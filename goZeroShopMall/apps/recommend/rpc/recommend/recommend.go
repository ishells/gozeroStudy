package main

import (
	"flag"
	"fmt"

	"goZeroShopMall/apps/recommend/rpc/recommend/internal/config"
	"goZeroShopMall/apps/recommend/rpc/recommend/internal/server"
	"goZeroShopMall/apps/recommend/rpc/recommend/internal/svc"
	"goZeroShopMall/apps/recommend/rpc/recommend/recommend"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/recommend.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		recommend.RegisterRecommendServer(grpcServer, server.NewRecommendServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
