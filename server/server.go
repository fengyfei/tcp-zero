// Package server provides a generic TCP server implementation.
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

	hubMutex sync.Mutex
	Hub      interfaces.Hub
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

		session := newSession(conn)
		srv.Put(session)

		if srv.Protocol != nil {
			go srv.Protocol.Handler(session, srv.close)
		}
	}
}

// Close the server immediately.
func (srv *Server) Close() (err error) {
	srv.once.Do(func() {
		close(srv.close)
		srv.Destroy()
	})

	return nil
}

// Put a new connection to hub.
func (srv *Server) Put(session interfaces.Session) error {
	if srv.Hub == nil {
		return nil
	}

	srv.hubMutex.Lock()
	defer srv.hubMutex.Unlock()

	return srv.Hub.Put(session)
}

// Remove a connection from hub, not responsible for closing the connection.
func (srv *Server) Remove(session interfaces.Session) error {
	if srv.Hub == nil {
		return nil
	}

	srv.hubMutex.Lock()
	defer srv.hubMutex.Unlock()

	return srv.Hub.Remove(session)
}

// Destroy a hub.
func (srv *Server) Destroy() error {
	if srv.Hub == nil {
		return nil
	}

	srv.hubMutex.Lock()
	defer srv.hubMutex.Unlock()

	return srv.Hub.Destroy()
}
