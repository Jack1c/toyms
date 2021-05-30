package server

import (
	"google.golang.org/grpc"
	"net"
)

type RpcServer struct {
	Address  string
	Register func(server *grpc.Server)
}

func (rs RpcServer) Start() error {
	listen, err := net.Listen("tcp", rs.Address)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	rs.Register(server)
	err = server.Serve(listen)
	return err
}
