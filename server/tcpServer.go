package server

import (
	log "github.com/sirupsen/logrus"
	"goRedis/server/connectionHandler"
	"net"
	"strconv"
)

type TCPServer struct {
	port      int
	isRunning bool
	handler   connectionHandler.ConnectionHandler
}

func NewTCPServer(port int, handler connectionHandler.ConnectionHandler) *TCPServer {
	return &TCPServer{port, false, handler}
}

func (ts *TCPServer) StartListen(params ...interface{}) {
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(ts.port))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer l.Close()
	log.Info("Starting To Listen on port: " + strconv.Itoa(ts.port))

	ts.isRunning = true
	for ts.isRunning {
		c, err := l.Accept()
		if err != nil {
			log.Error(err)
			continue
		}

		go ts.HandleConnection(c)
	}
}

func (ts *TCPServer) StopListen() {
	ts.isRunning = false
}

func (ts *TCPServer) HandleConnection(conn net.Conn) {
	if ts.handler == nil {
		panic("There is no connection handler! Shutting down...")
	}

	ts.handler.HandleConnection(conn)

	err := ts.handler.HandleConnection(conn)
	if err != nil {
		log.Error(conn.RemoteAddr(), err)
		return
	}
}
