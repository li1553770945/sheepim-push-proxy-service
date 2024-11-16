// Code generated by Kitex v0.7.2. DO NOT EDIT.
package projectservice

import (
	server "github.com/cloudwego/kitex/server"
	project "github.com/li1553770945/sheepim-online-service/kitex_gen/project"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler project.ProjectService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
