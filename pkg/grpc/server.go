package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	// Is GRPC Connected?
	IsConnected bool

	Connected chan int

	// GRPC Connection
	GRPCConn     *grpc.Server
	GRPCListener *net.Listener
}

// NewServer ...
func NewServer(connParams *Connection) (*Server, error) {

	serverConn := Server{
		IsConnected: false,
	}

	// Initialize GRPC Server
	err := serverConn.StartServer(connParams)
	if err != nil {
		return nil, err
	}

	return &serverConn, nil
}

// StartServer starts GRPC server
func (s *Server) StartServer(connParams *Connection) error {

	// Start GRPC server
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", connParams.IP, connParams.Port))
	if err != nil {
		return err
	}

	// Store listener
	s.GRPCListener = &ln

	srv := grpc.NewServer()

	/*
		log.Infof("gRPC Server: Serving at %s", fmt.Sprintf("%s:%d", connParams.IP, connParams.Port))
		if err := srv.Serve(ln); err != nil {
			return err
		}
	*/
	go srv.Serve(ln)

	s.GRPCConn = srv

	s.Connected <- StatusReady

	return nil
}
