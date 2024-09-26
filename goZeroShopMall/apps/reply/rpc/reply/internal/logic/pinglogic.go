package logic

import (
	"context"

	"goZeroShopMall/apps/reply/rpc/reply/internal/svc"
	"goZeroShopMall/apps/reply/rpc/reply/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *reply.Request) (*reply.Response, error) {
	// todo: add your logic here and delete this line

	return &reply.Response{}, nil
}
