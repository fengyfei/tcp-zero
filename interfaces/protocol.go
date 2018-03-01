// Package interfaces defines minimum behaviour of a TCP server.
package interfaces

// Protocol is a abstract handler for dealing with a tcp stream.
type Protocol interface {
	Handler(session Session, close <-chan struct{})
}

// ProtocolHandler wraps a common function as a Protocol.
type ProtocolHandler func(session Session, close <-chan struct{})

// Handler is a Protocal implementation.
func (f ProtocolHandler) Handler(session Session, close <-chan struct{}) {
	f(session, close)
}
