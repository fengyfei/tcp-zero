// Package interfaces provides a generic TCP server implementation.
package interfaces

import (
	"net"
)

// Session wraps common operations on a connection.
type Session interface {
	// Conn returns the underlying net.Conn.
	Conn() net.Conn

	// Send a message.
	Send(msg Message) error

	// Close the session.
	Close() error
}
