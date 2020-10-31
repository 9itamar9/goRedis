package connectionHandler

import (
	"bufio"
	"goRedis/parser"
	"net"
)

type RespHandler struct {
}

func (rh *RespHandler) HandleConnection(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	command := parser.ParseMessage(reader)
	return nil
}
