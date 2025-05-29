package user

import (
	"context"
	v1 "go-goframe-grpc/api/user/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	res = &v1.CreateUserRes{
		Code: 0,
		Msg:  "ok",
	}
	return res, nil
}
