package main

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	v1 "go-goframe-grpc/api/user/v1"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	//cmd.Main.Run(gctx.GetInitCtx())

	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	var (
		ctx    = gctx.New()
		conn   = grpcx.Client.MustNewGrpcClientConn("demo")
		client = v1.NewUserClient(conn)
	)
	res, err := client.CreateUser(ctx, &v1.CreateUserReq{
		Name: "张三",
		Sex:  "male",
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, "Response:", res.Msg)
}
