package connectionHandler

import "net"

type ConnectionHandler interface {
	HandleConnection(conn net.Conn) error
}