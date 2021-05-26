package main

import "google.golang.org/grpc"

func main() {
	conn, err := grpc.Dial("127.0.0.1:9988")
	if err != nil {
		panic(err)
	}

}
