package server

import "net"

type Server interface {
	StartListen(params ...interface{})
	StopListen()
	HandleConnection(conn net.Conn)
}