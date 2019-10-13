package main

import (
	"fmt"
	"sample-grpc-boilerplate/pkg/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Start GRPC client connection
	grpcServer, err := grpc.NewServer(&grpc.Connection{
		IP:   "127.0.0.1",
		Port: 8000,
	})

	if err != nil {
		log.Errorf("err: %v\n", err.Error())
		return
	}

	for {
		select {
		case grpcStatus := <-grpcServer.Connected:
			fmt.Println(grpcStatus)

		}
	}

}
