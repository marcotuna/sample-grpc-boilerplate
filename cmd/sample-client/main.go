package main

import "sample-grpc-boilerplate/pkg/grpc"

func main() {

	// Start GRPC client connection
	grpc.NewClient(&grpc.Connection{
		IP:   "127.0.0.1",
		Port: 8000,
	})

}
