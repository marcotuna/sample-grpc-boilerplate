package grpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Client ...
type Client struct {
	// Is GRPC connected?
	IsConnected bool
	Connected   chan connectivity.State
	// GRPC connection
	GRPCConn *grpc.ClientConn
}

// NewClient ...
func NewClient(connParams *Connection) (*Client, error) {

	clientConn := Client{
		IsConnected: false,
		Connected:   make(chan connectivity.State),
	}

	// Initialize GRPC Client
	err := clientConn.StartClient(connParams)
	if err != nil {
		return nil, err
	}

	return &clientConn, nil
}

// StartClient creates a new GRPC Client
func (c *Client) StartClient(connParams *Connection) error {

	cc, err := grpc.Dial(
		fmt.Sprintf("%s:%d", connParams.IP, connParams.Port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if err != nil {
		return err
	}

	// Store GRPC connection
	c.GRPCConn = cc

	// Update connection status to true
	c.IsConnected = true

	// Update channel status to true
	go func() {
		c.Connected <- StatusReady
	}()

	// Start GRPC health checking
	go c.grpcHealth(cc)

	return nil
}

// Shutdown closes client connection
func (c *Client) Shutdown() {
	c.GRPCConn.Close()
}

func (c *Client) grpcHealth(cc *grpc.ClientConn) {

	// Create ticker and run every 500 ms
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Run ticker immediatly
	for ; true; <-ticker.C {
		switch cc.GetState() {
		case connectivity.Connecting, connectivity.Idle, connectivity.Shutdown, connectivity.TransientFailure:
			// Connection is unavailable
			// Emit event with status false
			if c.IsConnected {
				c.Connected <- cc.GetState()
			}

			// Set status to false
			c.IsConnected = false
		case connectivity.Ready:

			// Connection is ready
			// Emit event with status true
			if !c.IsConnected {
				c.Connected <- cc.GetState()
			}

			// Set status to true
			c.IsConnected = true
		}

	}
}
