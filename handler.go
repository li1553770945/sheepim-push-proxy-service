package main

import (
	"context"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/container"
	push_proxy "github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
)

// PushProxyServiceImpl implements the last service interface defined in the IDL.
type PushProxyServiceImpl struct{}

// PushMessage implements the PushProxyServiceImpl interface.
func (s *PushProxyServiceImpl) PushMessage(ctx context.Context, req *push_proxy.PushMessageReq) (resp *push_proxy.PushMessageResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.PushProxyService.PushMessage(ctx, req)
	return
}
