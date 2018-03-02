// A simple TCP server.
package main

import (
	"fmt"
	"bufio"

	"github.com/fengyfei/tcp-zero/interfaces"
	"github.com/fengyfei/tcp-zero/server"
)

func main() {
	srv := server.NewServer(":9573", &Handler{})

	if err := srv.ListenAndServe(); err != nil {
		srv.Close()
	}
}

type Handler struct {}

func (h *Handler) Handler(session interfaces.Session, close <-chan struct{}) {
	conn := session.Conn()
	reader := bufio.NewReader(conn)

	msg := server.NewMsg("welcome to tcp-zero \n")
	session.Send(msg)

	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			break
		}

		msg := server.NewMsg(fmt.Sprintf("you said: %s \n", line))
		session.Send(msg)
		fmt.Print("receive ", line)
	}

}
