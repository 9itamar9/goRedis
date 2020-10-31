package connectionHandler

import (
	"net"
)

type RespHandler struct {
}

func (rh *RespHandler) HandleConnection(conn net.Conn) error {
	panic("implement me")
}

