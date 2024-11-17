package article

import (
	"context"
	"learn_go/webook/internal/domain"
	"learn_go/webook/internal/repository/dao"
)

type ArticleRepository interface {
	Create(ctx context.Context, article domain.Article) (int64, error)
	Update(ctx context.Context, article domain.Article) error

	Sync(ctx context.Context, article domain.Article) (int64, error)

	// SyncV1 在repository层同步数据
	SyncV1(ctx context.Context, article domain.Article) (int64, error)
}

func NewArticleRepository(articleDao dao.ArticleDao) ArticleRepository {
	return &articleRepository{articleDao: articleDao}
}

type articleRepository struct {
	articleDao dao.ArticleDao

	articleAuthorDao dao.ArticleAuthorDao
	articleReaderDao dao.ArticleReaderDao
}

func (repo *articleRepository) Sync(ctx context.Context, article domain.Article) (int64, error) {
	// 如果是同个库，可以把数据的同步交给dao层。
	return repo.articleDao.Sync(ctx, repo.toEntity(article))
}

func (repo *articleRepository) SyncV1(ctx context.Context, article domain.Article) (int64, error) {
	// 数据可能是俩个存储源（比如不同库），就由不同存储源所对应的dao来处理。
	var (
		id  = article.ID
		err error
	)
	if article.ID > 0 {
		err = repo.articleAuthorDao.UpdateByID(ctx, repo.toEntity(article))
	} else {
		id, err = repo.articleAuthorDao.Insert(ctx, repo.toEntity(article))
	}
	if err != nil {
		return 0, err
	}
	article.ID = id
	err = repo.articleReaderDao.Upsert(ctx, repo.toEntity(article))
	return id, err
}

func (repo *articleRepository) Create(ctx context.Context, article domain.Article) (int64, error) {
	return repo.articleDao.Insert(ctx, repo.toEntity(article))
}

func (repo *articleRepository) Update(ctx context.Context, article domain.Article) error {
	// tip:
	// 用户只能更新自己的帖子。先查询再判定的性能不好，因为多了一次查询。
	// 正常用户是不会出现更新其他作者的帖子的，因为可以在更新时进行条件限制。
	return repo.articleDao.UpdateByID(ctx, repo.toEntity(article))
}

func (repo *articleRepository) toEntity(article domain.Article) dao.Article {
	return dao.Article{
		Title:    article.Title,
		Content:  article.Content,
		ID:       article.ID,
		AuthorID: article.Author.ID,
		Status:   article.Status.ToInt8(),
	}
}
