// Package interfaces provides a generic TCP server implementation.
package interfaces

import (
	"net"
)

// Session wraps common operations on a connection.
type Session interface {
	// Conn returns the underlying net.Conn.
	Conn() net.Conn

	// Put a message to queue.
	Put(msg Message) bool

	// Send a message.
	Send() bool

	// Close the session.
	Close() error
}
