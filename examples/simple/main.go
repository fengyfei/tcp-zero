/*
 * Revision History:
 *     Initial: 2018/02/28        Feng Yifei
 */

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
