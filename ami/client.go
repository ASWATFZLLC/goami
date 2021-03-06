package ami

// Client defines an interface to client socket.
type Client interface {
	// Connected returns the client status.
	Connected() bool

	// Close closes the client connection.
	Close() error

	// Send sends data from client to server.
	Send(message string) error

	// Recv receives a string from server.
	Recv() (string, error)
}
