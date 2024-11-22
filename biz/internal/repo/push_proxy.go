package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
	"github.com/segmentio/kafka-go"
)

func (r *Repository) PushMessage(ctx context.Context, messageObj *push_proxy.PushMessageReq) error {
	messageBytes, err := json.Marshal(messageObj)
	if err != nil {
		return errors.New(fmt.Sprintf("消息序列化失败 %v", err))
	}

	// 创建 Kafka 消息
	msg := kafka.Message{
		Key:   []byte(messageObj.ClientId), // 使用 ClientId 作为消息的 Key
		Value: messageBytes,                // 消息体为序列化后的 JSON
	}

	// 使用 Kafka Producer 发送消息
	err = r.KafkaClient.Producer.WriteMessages(ctx, msg)
	if err != nil {
		return errors.New(fmt.Sprintf("消息队列消息发送失败:%v", err))
	}
	klog.CtxInfof(ctx, "放入消息成功:%s", messageObj.Message)

	return nil
}
