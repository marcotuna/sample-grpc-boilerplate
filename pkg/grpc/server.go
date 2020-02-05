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
		Connected:   make(chan int),
	}

	// Initialize GRPC Server
	err := serverConn.Start(connParams)
	if err != nil {
		return nil, err
	}

	return &serverConn, nil
}

// Start starts GRPC server
func (s *Server) Start(connParams *Connection) error {

	// Start GRPC server
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", connParams.IP, connParams.Port))
	if err != nil {
		return err
	}

	// Store listener
	s.GRPCListener = &ln

	srv := grpc.NewServer()

	// Start serving requests
	go srv.Serve(ln)
	s.GRPCConn = srv

	// Notify connection ready
	go func() {
		s.Connected <- StatusReady
	}()

	return nil
}

// Shutdown closes the running GRPC Server
func (s *Server) Shutdown() {
	// Shutdown GRPC
	s.GRPCConn.GracefulStop()
}
