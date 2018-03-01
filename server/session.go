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

func (s *session) Send(msg interfaces.Message) bool {
	return s.queue.Put(msg)
}

func (s *session) Close() error {
	s.queue.Close()

	return s.conn.Close()
}
