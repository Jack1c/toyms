package server

import (
	"google.golang.org/grpc"
	"net"
)

type RpcServer struct {
	address  string
	register func(server *grpc.Server)
}

func Start(g RpcServer) error {
	listen, err := net.Listen("tcp", g.address)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	g.register(server)
	err = server.Serve(listen)
	return err
}
