package main

import (
	"fmt"
	"sample-grpc-boilerplate/pkg/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("GRPC Server")

	grpcConnParams := &grpc.Connection{
		IP:   "127.0.0.1",
		Port: 8000,
	}

	// Start GRPC client connection
	grpcServer, err := grpc.NewServer(&grpc.Connection{
		IP:   grpcConnParams.IP,
		Port: grpcConnParams.Port,
	})

	defer grpcServer.Shutdown()

	if err != nil {
		log.Errorf("err: %v\n", err.Error())
		return
	}

	// Notify about status changes on GRPC connection
	for {
		select {
		case grpcStatus := <-grpcServer.Connected:
			switch grpcStatus {
			case grpc.StatusConnecting:
				fmt.Printf("gRPC Server: Connecting\n")
			case grpc.StatusReady:
				fmt.Printf("gRPC Server: Serving at %s\n", fmt.Sprintf("%s:%d", grpcConnParams.IP, grpcConnParams.Port))
			case grpc.StatusShutdown:
				fmt.Printf("gRPC Server: Shutting down...\n")
			case grpc.StatusTransientFailure, grpc.StatusIdle, grpc.StatusInvalid:
				fmt.Printf("gRPC Server: Error\n")
			}
		}
	}

}
