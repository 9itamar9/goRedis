package server

import "net"

type Server interface {
	StartListen()
	StopListen()
	HandleConnection(conn net.Conn)
}