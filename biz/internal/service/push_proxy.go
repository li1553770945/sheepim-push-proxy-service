package service

import (
	"context"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/constant"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
)

func (s PushProxyService) PushMessage(ctx context.Context, req *push_proxy.PushMessageReq) (resp *push_proxy.PushMessageResp, err error) {
	err = s.Repo.PushMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return &push_proxy.PushMessageResp{
		BaseResp: &base.BaseResp{Code: constant.Success},
	}, nil
}
