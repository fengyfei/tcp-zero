/*
 * Revision History:
 *     Initial: 2018/02/28        Feng Yifei
 */

package server

import (
	"net"
	"sync"
	"time"

	"github.com/fengyfei/tcp-zero/interfaces"
)

// Server represents a general TCP server.
type Server struct {
	Addr     string
	Protocol interfaces.Protocol
	listener net.Listener
	close    chan struct{}
	once     sync.Once
}

// NewServer creates a non-TLS TCP server.
func NewServer(addr string, protocol interfaces.Protocol) *Server {
	return &Server{
		Addr:     addr,
		Protocol: protocol,
		close:    make(chan struct{}),
	}
}

// ListenAndServe listens on a address and serves the incomming connections.
func (srv *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}

	return srv.Serve(l)
}

// Serve on the given listener.
func (srv *Server) Serve(l net.Listener) error {
	srv.listener = l
	defer func() {
		srv.listener.Close()
	}()

	for {
		var (
			conn net.Conn
			err  error
		)

		conn, err = srv.listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			return err
		}

		if srv.Protocol != nil {
			go srv.Protocol.Handler(conn, srv.close)
		}
	}
}

// Close the server immediately.
func (srv *Server) Close() (err error) {
	srv.once.Do(func() {
		close(srv.close)
	})

	return nil
}
