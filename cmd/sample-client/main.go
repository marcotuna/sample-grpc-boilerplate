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

	defer grpcClient.Shutdown()

	if err != nil {
		log.Errorf("err: %v\n", err.Error())
		return
	}

	// Notify about status changes on GRPC connection
	for {
		select {
		case grpcStatus := <-grpcClient.Connected:
			switch grpcStatus {
			case grpc.StatusConnecting:
				fmt.Printf("gRPC Client: Connecting\n")
			case grpc.StatusReady:
				fmt.Printf("gRPC Client: Serving at %s\n", fmt.Sprintf("%s:%d", grpcConnParams.IP, grpcConnParams.Port))
			case grpc.StatusShutdown:
				fmt.Printf("gRPC Client: Shutting down...\n")
			case grpc.StatusTransientFailure, grpc.StatusIdle, grpc.StatusInvalid:
				fmt.Printf("gRPC Client: Error\n")
			}

		}
	}

}
