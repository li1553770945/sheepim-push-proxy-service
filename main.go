// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/container"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy/pushproxyservice"
	"net"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	container.InitGlobalContainer(env)
	App := container.GetGlobalContainer()

	serviceName := App.Config.ServerConfig.ServiceName

	defer func(container *container.Container) {
		p := container.Trace.Provider
		ctx := context.Background()
		err := p.Shutdown(ctx)
		if err != nil {
			klog.Errorf("server stopped with error:%s", err)
		}

		if container.KafkaClient.Producer != nil {
			err := container.KafkaClient.Producer.Close()
			if err != nil {
				klog.Errorf("kafka生产者关闭失败%s", err)
			}
		}

		if container.KafkaClient.Consumer != nil {
			err := container.KafkaClient.Consumer.Close()
			if err != nil {
				klog.Errorf("kafka消费者关闭失败%s", err)
			}
		}

	}(App)

	addr, err := net.ResolveTCPAddr("tcp", App.Config.ServerConfig.ListenAddress)
	if err != nil {
		panic("设置监听地址出错")
	}

	r, err := etcd.NewEtcdRegistry(App.Config.EtcdConfig.Endpoint) // r should not be reused.
	if err != nil {
		panic(err)
	}
	svr := pushproxyservice.NewServer(
		new(PushProxyServiceImpl),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("服务启动失败:", err)
	}
}
