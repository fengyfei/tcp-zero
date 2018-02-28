// A simple TCP server.
package main

import (
	"github.com/fengyfei/tcp-zero/server"
)

func main() {
	srv := server.NewServer(":9573", nil)

	if err := srv.ListenAndServe(); err != nil {
		srv.Close()
	}
}
