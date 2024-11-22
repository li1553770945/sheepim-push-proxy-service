package container

import (
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/config"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/kafka"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/log"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/internal/service"
	"sync"
)

type Container struct {
	Trace            *trace.TraceStruct
	Logger           *log.TraceLogger
	KafkaClient      *kafka.KafkaClient
	Config           *config.Config
	PushProxyService service.IPushProxyService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(env string) {
	once.Do(func() {
		APP = GetContainer(env)
	})
}

func NewContainer(config *config.Config,
	logger *log.TraceLogger,
	traceStruct *trace.TraceStruct,
	kafkaClient *kafka.KafkaClient,
	pushProxyService service.IPushProxyService,
) *Container {
	return &Container{
		Config:           config,
		Logger:           logger,
		KafkaClient:      kafkaClient,
		PushProxyService: pushProxyService,
		Trace:            traceStruct,
	}

}
