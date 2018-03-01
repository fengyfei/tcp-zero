// Package interfaces provides a generic TCP server implementation.
package interfaces

import (
	"net"
)

// Session wraps common operations on a connection.
type Session interface {
	Conn() net.Conn
	Send(msg Message) error
	Close() error
}
