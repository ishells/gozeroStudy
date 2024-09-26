package logic

import (
	"context"

	"goZeroShopMall/apps/recommend/rpc/recommend/internal/svc"
	"goZeroShopMall/apps/recommend/rpc/recommend/recommend"

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

func (l *PingLogic) Ping(in *recommend.Request) (*recommend.Response, error) {
	// todo: add your logic here and delete this line

	return &recommend.Response{}, nil
}
