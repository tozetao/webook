package article

import (
	"context"
	"github.com/IBM/sarama"
	"learn_go/webook/interaction/event"
	"learn_go/webook/interaction/repository"
	"learn_go/webook/pkg/logger"
	"learn_go/webook/pkg/saramax"
)

// BatchReadEventConsumer 消费者：消费生产者投递的各种事件（消息）
type BatchReadEventConsumer struct {
	client          sarama.Client
	interactionRepo repository.InteractionRepository
	l               logger.LoggerV2
}

func NewBatchReadEventConsumer(client sarama.Client, interactionRepo repository.InteractionRepository,
	l logger.LoggerV2) *BatchReadEventConsumer {
	return &BatchReadEventConsumer{
		interactionRepo: interactionRepo,
		client:          client,
		l:               l,
	}
}

func (c *BatchReadEventConsumer) Start() error {
	// 创建一个消费组用于消费读取文章事件
	consumer, err := sarama.NewConsumerGroupFromClient(groupInteraction, c.client)
	if err != nil {
		return err
	}
	go func() {
		err := consumer.Consume(context.Background(),
			[]string{event.TopicArticleReadEvent},
			saramax.NewBatchHandler[event.ReadEvent](c.l, c.Consume))
		if err != nil {
			c.l.Error("Article消费者退出", logger.Error(err))
		}
	}()
	return nil
}

// Consume 消费文章读取时间
func (c *BatchReadEventConsumer) Consume(messages []*sarama.ConsumerMessage, events []event.ReadEvent) error {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	//defer cancel()
	ctx := context.Background()

	bizs := make([]string, 0, len(events))
	bizIDs := make([]int64, 0, len(events))
	for _, evt := range events {
		bizs = append(bizs, "article")
		bizIDs = append(bizIDs, evt.ArticleID)
	}
	return c.interactionRepo.BatchIncrReadCnt(ctx, bizs, bizIDs)
}
