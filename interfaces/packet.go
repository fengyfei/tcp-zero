// Package interfaces provides a generic TCP server implementation.
package interfaces

// Encoder is the interface implemented by types to convert to network-endian.
type Encoder interface {
	Encode() ([]byte, error)
}

// Decoder is the interface implemented by types to convert to host-endian.
type Decoder interface {
	Decode([]byte) error
}

// Message is the interface implemented by types transferred on network.
type Message interface {
	Encoder
	Decoder
}
