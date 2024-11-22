//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/config"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/kafka"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/log"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/internal/service"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.GetConfig,
		log.InitLog,
		trace.InitTrace,
		kafka.NewKafkaClient,
		//repo
		repo.NewRepository,

		//service
		service.NewPushProxyService,

		NewContainer,
	))
}
