package grpc

const (
	// StatusConnecting ...
	StatusConnecting = 1
	// StatusReady ...
	StatusReady = 2
	// StatusTransientFailure ...
	StatusTransientFailure = 3
	// StatusIdle ...
	StatusIdle = 4
	// StatusShutdown ...
	StatusShutdown = 5
	// StatusInvalid ...
	StatusInvalid = 6
)

// Connection ...
type Connection struct {
	IP       string
	Port     int
	Username string
	Password string
}
