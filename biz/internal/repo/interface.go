package repo

import (
	"context"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/kafka"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
)

type IRepository interface {
	PushMessage(ctx context.Context, messageObj *push_proxy.PushMessageReq) error
}

type Repository struct {
	KafkaClient *kafka.KafkaClient
}

func NewRepository(kafkaClient *kafka.KafkaClient) IRepository {
	return &Repository{
		KafkaClient: kafkaClient,
	}
}
