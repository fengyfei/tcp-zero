// Package interfaces defines minimum behaviour of a TCP server.
package interfaces

import (
	"net"
)

// Hub provides group methods on connections.
type Hub interface {
	Put(net.Conn) error
	Remove(net.Conn) error
	Destroy() error
}
