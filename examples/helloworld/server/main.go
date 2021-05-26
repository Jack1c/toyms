package main

import (
	"context"
	"github.com/Jack1c/toyms/examples/helloworld/pb"
	"google.golang.org/grpc"
	"net"
)

type HelloService struct{}

func (hs *HelloService) SayHello(ctx context.Context, req *pb.HelloReq) (rsp *pb.HelloRsp, err error) {
	return &pb.HelloRsp{Rt: "Hello:" + req.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9988")
	if err != nil {
		panic(err)
	}

	grpcs := grpc.NewServer()
	pb.RegisterHelloWordServer(grpcs, &HelloService{})
	err = grpcs.Serve(listen)
	if err != nil {
		panic(err)
	}
}
