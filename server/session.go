// Package server provides a generic TCP server implementation.
package server

import (
	"net"

	"github.com/fengyfei/tcp-zero/interfaces"
)

type session struct {
	conn net.Conn
}

func newSession(conn net.Conn) interfaces.Session {
	return &session{
		conn: conn,
	}
}

func (s *session) Conn() net.Conn {
	return s.conn
}

func (s *session) Send(msg interfaces.Message) error {
	return nil
}

func (s *session) Close() error {
	return nil
}
