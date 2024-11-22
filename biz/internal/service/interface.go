package service

import (
	"context"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
)

type PushProxyService struct {
	Repo repo.IRepository
}

type IPushProxyService interface {
	PushMessage(ctx context.Context, req *push_proxy.PushMessageReq) (resp *push_proxy.PushMessageResp, err error)
}

func NewPushProxyService(repo repo.IRepository) IPushProxyService {
	return &PushProxyService{
		Repo: repo,
	}
}
