package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goZeroShopMall/apps/app/api/internal/config"

	"goZeroShopMall/apps/order/rpc/orderclient"
	"goZeroShopMall/apps/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
