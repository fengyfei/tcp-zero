// Package interfaces defines minimum behaviour of a TCP server.
package interfaces

// Hub provides group methods on connections.
type Hub interface {
	Put(Session) error
	Remove(Session) error
	Destroy() error
}
