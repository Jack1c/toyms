package main

import (
	"context"
	"fmt"
	"github.com/Jack1c/toyms/examples/helloworld/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9988", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	rsp, err := pb.NewHelloWordClient(conn).SayHello(context.Background(), &pb.HelloReq{Name: "grpc"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp.Rt)
}
