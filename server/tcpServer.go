package server

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"goRedis/parser"
	"net"
	"strconv"
)

type TCPServer struct {
	port      int
	isRunning bool
	pr        parser.Parser
}

func NewTCPServer(port int, pr parser.Parser) *TCPServer {
	return &TCPServer{port, false, pr}
}

func (ts *TCPServer) StartListen() {
	listener, err := net.Listen("tcp4", ":"+strconv.Itoa(ts.port))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer listener.Close()
	log.Info("Starting To Listen on port: " + strconv.Itoa(ts.port))

	ts.isRunning = true
	for ts.isRunning {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err)
			continue
		}

		go ts.HandleConnection(conn)
	}
}

func (ts *TCPServer) StopListen() {
	ts.isRunning = false
}

func (ts *TCPServer) HandleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	command, err := ts.pr.ParseMessage(reader)

	if err != nil {
		log.Error(err)
		return
	} else if _, ok := command.([]interface{}); !ok {
		log.Error(fmt.Sprintf("Not a known command structure: %v", command))
		return
	}

	log.Info(fmt.Sprintf("got %v from %v", command, conn.RemoteAddr()))
}
