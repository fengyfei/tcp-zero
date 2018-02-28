/*
 * Revision History:
 *     Initial: 2018/02/28        Feng Yifei
 */

package interfaces

import (
	"net"
)

// Protocol is a abstract handler for dealing with a tcp stream.
type Protocol interface {
	Handler(conn net.Conn, close <-chan struct{})
}
