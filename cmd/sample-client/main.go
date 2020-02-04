package main

import (
	"fmt"
	"sample-grpc-boilerplate/pkg/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("GRPC Client")

	grpcConnParams := &grpc.Connection{
		IP:   "127.0.0.1",
		Port: 8000,
	}

	// Start GRPC client connection
	grpcClient, err := grpc.NewClient(&grpc.Connection{
		IP:   grpcConnParams.IP,
		Port: grpcConnParams.Port,
	})

	if err != nil {
		log.Errorf("err: %v\n", err.Error())
		return
	}

	// Notify about status changes on GRPC connection
	for {
		select {
		case grpcStatus := <-grpcClient.Connected:
			fmt.Println(grpcStatus)

			switch grpcStatus {
			case 1:
				fmt.Printf("gRPC Client: Connecting\n")
			case 2:
				fmt.Printf("gRPC Client: Serving at %s\n", fmt.Sprintf("%s:%d", grpcConnParams.IP, grpcConnParams.Port))
			case 5:
				fmt.Printf("gRPC Client: Shutting down...\n")
			case 3, 4, 6:
				fmt.Printf("gRPC Client: Error\n")
			}

		}
	}

}
