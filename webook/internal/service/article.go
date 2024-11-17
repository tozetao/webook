package service

import (
	"context"
	"errors"
	"learn_go/webook/internal/domain"
	"learn_go/webook/internal/repository/article"
	"learn_go/webook/pkg/logger"
)

/*
Publish的数据同步。

发布这种操作会有3种情况：
1. 新建Article并发表：插入制作库 => 插入线上库
2. 编辑Article后首次发表：更新制作库 => 插入线上库
3. 编辑Article并发表：更新制作库 => 更新线上库

可以发现制作库的数据需要同步到线上库，在代码层面上可以着由划分：
ArticleAuthorService
	Save(domain.Article)
ArticleReaderService
	Save(domain.Article)
偏向微服务的划分。根据线上库、制作库划分为俩个不同的服务，最后在web层这俩种服务。


在Repository层分为ArticleAuthorRepository、ArticleReaderRepository，
在Service层聚合这俩个仓库，实现发布。

将数据同步的操作封装到ArticleRepository种。
ArticleRepository
	Save(domain.Article)


如果需要本地事务，必然是同库不同表，这时候可以在Repository或Dao层中处理。如果非本地事务，可以在任意层面上处理，

note: 如果俩个数据源的结构相似度较高，代码相似度较高，个人偏向将同步数据的操作交给Repository来处理。


*/

type ArticleService interface {
	Save(ctx context.Context, article domain.Article) (int64, error)

	Publish(ctx context.Context, article domain.Article) (int64, error)

	PublishV1(ctx context.Context, article domain.Article) (int64, error)
}

type articleService struct {
	log logger.LoggerV2

	articleRepo article.ArticleRepository

	// 与ArticleRepository互斥
	articleAuthorRepo article.AuthorRepository
	articleReaderRepo article.ReaderRepository
}

func (svc *articleService) Publish(ctx context.Context, article domain.Article) (int64, error) {
	article.Status = domain.ArticleStatusPublished
	return svc.articleRepo.Sync(ctx, article)
}

func (svc *articleService) PublishV1(ctx context.Context, article domain.Article) (int64, error) {
	var (
		id  = article.ID
		err error
	)
	if article.ID > 0 {
		err = svc.articleAuthorRepo.Update(ctx, article)
	} else {
		id, err = svc.articleAuthorRepo.Create(ctx, article)
	}
	article.ID = id
	if err != nil {
		return 0, err
	}

	/**
	新建文章 => 发布文章，这俩个操作并不是原子性的，在新建或者发布都可能发生失败。

	能否使用事务？
		按照Repository的接口设计，是无法使用事务的。因为无法保证Author、Reader是同个库，或者同个存储源。
		并且对于是纯关系型数据库的事务，开启事务和提交事务不应该放在Service层来控制。

	解决方案?
		1. 重试，尽可能让其成功。注：多个数据存储源很难保持数据强一致性，都是尽可能的保证数据的最终趋于一致。系统的可用性也是一个道理，很难做到100%。
		2. 在重试后通过cancel、消息队列去处理。
		3. 更上层一点，发布文章这个业务可以通过监听新建文章的binlog，用canal来实现。

	总结来说一般业务系统不要做的很复杂，重试就够了。
	*/
	for i := 0; i < 3; i++ {
		id, err = svc.articleReaderRepo.Save(ctx, article)
		if err == nil {
			return id, nil
		}
		svc.log.Error("线上库发布失败", logger.Int64("article id", article.ID), logger.Error(err))
	}
	svc.log.Error("线上库发布失败，重试次数耗尽", logger.Int64("article id", article.ID), logger.Error(err))
	return 0, errors.New("failed to publish")
}

func NewArticleService(
	articleRepo article.ArticleRepository,
	articleAuthorRepo article.AuthorRepository,
	articleReaderRepo article.ReaderRepository,
	log logger.LoggerV2) ArticleService {
	return &articleService{
		log:               log,
		articleRepo:       articleRepo,
		articleAuthorRepo: articleAuthorRepo,
		articleReaderRepo: articleReaderRepo,
	}
}

// Save status = unpublish，
func (svc *articleService) Save(ctx context.Context, article domain.Article) (int64, error) {
	article.Status = domain.ArticleStatusUnpublished
	if article.ID > 0 {
		err := svc.articleRepo.Update(ctx, article)
		return article.ID, err
	}
	return svc.articleRepo.Create(ctx, article)
}