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

	// Start serving requests
	go srv.Serve(ln)
	s.GRPCConn = srv

	// Notify connection ready
	go func() {
		s.Connected <- StatusReady
	}()

	return nil
}

// ShutdownServer closes the running GRPC Server
func (s *Server) ShutdownServer() {
	// Shutdown GRPC
	s.GRPCConn.GracefulStop()
}
