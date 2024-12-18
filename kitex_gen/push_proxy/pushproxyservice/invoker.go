// Code generated by Kitex v0.7.2. DO NOT EDIT.

package pushproxyservice

import (
	server "github.com/cloudwego/kitex/server"
	push_proxy "github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler push_proxy.PushProxyService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
