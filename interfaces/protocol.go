// Package interfaces defines minimum behaviour of a TCP server.
package interfaces

import (
	"net"
)

// Protocol is a abstract handler for dealing with a tcp stream.
type Protocol interface {
	Handler(conn net.Conn, close <-chan struct{})
}

// ProtocolHandler wraps a common function as a Protocol.
type ProtocolHandler func(conn net.Conn, close <-chan struct{})

// Handler is a Protocal implementation.
func (f ProtocolHandler) Handler(conn net.Conn, close <-chan struct{}) {
	f(conn, close)
}
