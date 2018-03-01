// Package server provides a generic TCP server implementation.
package server

import (
	"sync"

	"github.com/fengyfei/tcp-zero/interfaces"
)

type simpleQueue struct {
	c chan interfaces.Message

	mu     sync.RWMutex
	closed bool

	closeCh chan bool
}

func newQueue(size int) interfaces.Queue {
	if size <= 0 {
		size = 16
	}

	return &simpleQueue{
		c:       make(chan interfaces.Message, size),
		closeCh: make(chan bool),
	}
}

func (q *simpleQueue) Put(msg interfaces.Message) bool {
	if q.isClosed() {
		return false
	}

	select {
	case q.c <- msg:
		return true
	case <-q.closeCh:
		return false
	}
}

func (q *simpleQueue) Get() (interfaces.Message, bool) {
	if q.isClosed() {
		return nil, false
	}

	select {
	case msg := <-q.c:
		return msg, true
	case <-q.closeCh:
		return nil, false
	default:
		return nil, false
	}
}

func (q *simpleQueue) Close() {
	q.mu.Lock()
	q.closed = true
	q.mu.Unlock()

	// Notify all the running routines.
	close(q.closeCh)
}

func (q *simpleQueue) Closed() bool {
	return q.isClosed()
}

func (q *simpleQueue) Wait() (interfaces.Message, bool) {
	if q.isClosed() {
		return nil, false
	}

	select {
	case msg := <-q.c:
		return msg, true
	case <-q.closeCh:
		return nil, false
	}
}

func (q *simpleQueue) isClosed() bool {
	q.mu.RLock()
	closed := q.closed
	q.mu.RUnlock()

	return closed
}
