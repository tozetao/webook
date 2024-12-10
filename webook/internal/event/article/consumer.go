package article

import (
	"context"
	"github.com/IBM/sarama"
	"learn_go/webook/internal/event"
	"learn_go/webook/internal/repository"
	"learn_go/webook/pkg/logger"
	saramax "learn_go/webook/pkg/sarama"
	"time"
)

// 阅读文章事件的消费者

// 定义group ID
const (
	groupInteraction = "group:interaction"
)

// Consumer 消费者：消费生产者投递的各种事件（消息）
type Consumer struct {
	client          sarama.Client
	interactionRepo repository.InteractionRepository
	l               logger.LoggerV2
}

func NewConsumer(client sarama.Client, interactionRepo repository.InteractionRepository,
	l logger.LoggerV2) event.Consumer {
	return &Consumer{
		interactionRepo: interactionRepo,
		client:          client,
		l:               l,
	}
}

func (c *Consumer) Start() error {
	// 创建一个消费组用于消费读取文章事件
	consumer, err := sarama.NewConsumerGroupFromClient(groupInteraction, c.client)
	if err != nil {
		return err
	}
	go func() {
		err := consumer.Consume(context.Background(),
			[]string{TopicReadEvent},
			saramax.NewHandler[ReadEvent](c.l, c.Consume))
		if err != nil {
			c.l.Error("Article消费者退出", logger.Error(err))
		}
	}()
	return nil
}

// Consume 消费文章读取时间
func (c *Consumer) Consume(event ReadEvent, topic string, partition int32, offset int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	return c.interactionRepo.IncrReadCnt(ctx, "article", event.ArticleID)
}
