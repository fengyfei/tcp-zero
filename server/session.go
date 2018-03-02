// Package server provides a generic TCP server implementation.
package server

import (
	"net"

	"github.com/fengyfei/tcp-zero/interfaces"
)

type session struct {
	conn  net.Conn
	queue interfaces.Queue
}

func newSession(conn net.Conn) interfaces.Session {
	return &session{
		conn:  conn,
		queue: newQueue(0),
	}
}

func (s *session) Conn() net.Conn {
	return s.conn
}

func (s *session) Put(msg interfaces.Message) bool {
	return s.queue.Put(msg)
}

func (s *session) Send() bool {
	msg, ok := s.queue.Wait()
	if !ok {
		return ok
	}

	b, err := msg.Encode()
	if err != nil {
		return false
	}

	_, err = s.conn.Write(b)
	if err != nil {
		return false
	}

	return true
}

func (s *session) Close() error {
	s.queue.Close()

	return s.conn.Close()
}
