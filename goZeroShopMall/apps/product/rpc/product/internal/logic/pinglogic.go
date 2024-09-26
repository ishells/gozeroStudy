package logic

import (
	"context"

	"goZeroShopMall/apps/product/rpc/product/internal/svc"
	"goZeroShopMall/apps/product/rpc/product/product"

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

func (l *PingLogic) Ping(in *product.Request) (*product.Response, error) {
	// todo: add your logic here and delete this line

	return &product.Response{}, nil
}
