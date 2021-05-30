package app

import (
	"context"
	"github.com/Jack1c/toyms/examples/helloworld/pb"
	"github.com/Jack1c/toyms/server"
	"google.golang.org/grpc"
	"testing"
)

type HelloService struct{}

func (hs *HelloService) SayHello(ctx context.Context, req *pb.HelloReq) (rsp *pb.HelloRsp, err error) {
	return &pb.HelloRsp{Rt: "Hello:" + req.Name}, nil
}

func TestApp(t *testing.T) {
	app := NewApp(Name("hello"), Versopn("v1.0"),
		Server(&server.RpcServer{
			Address: ":8899",
			Register: func(server *grpc.Server) {
				pb.RegisterHelloWordServer(server, &HelloService{})
			},
		}),
	)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
