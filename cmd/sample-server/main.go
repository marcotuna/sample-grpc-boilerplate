package main

import (
	"fmt"
	"sample-grpc-boilerplate/pkg/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {

	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{})

	logger.Info("gRPC Server")

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
		logger.Errorf("err: %v\n", err.Error())
		return
	}

	// Notify about status changes on GRPC connection
	for {
		select {
		case grpcStatus := <-grpcServer.Connected:
			switch grpcStatus {
			case grpc.StatusConnecting:
				logger.Info("gRPC Server: Connecting")
			case grpc.StatusReady:
				logger.Infof("gRPC Server: Serving at %s\n", fmt.Sprintf("%s:%d", grpcConnParams.IP, grpcConnParams.Port))
			case grpc.StatusShutdown:
				logger.Info("gRPC Server: Shutting down...")
			case grpc.StatusTransientFailure, grpc.StatusIdle, grpc.StatusInvalid:
				logger.Error("gRPC Server: Error")
			}
		}
	}

}
