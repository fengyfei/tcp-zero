// Package interfaces provides a generic TCP server implementation.
package interfaces

// Queue buffers messages should be sent on a connection.
type Queue interface {
	// Put a message in the queue, return false if queue is closed.
	Put(Message) bool

	// Get a message from the queue, return false if queue is empty or is closed.
	Get() (Message, bool)

	// Close the queue and discard all entries in the queue.
	Close()

	// Closed returns true if the queue is closed.
	Closed() bool

	// Wait for a message, returns false when the queue is closed.
	Wait() (Message, bool)
}
