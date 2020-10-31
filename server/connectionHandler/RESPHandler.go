package connectionHandler

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
)

const delimiter = "\r\n"

var convertToVal = map[byte]func(re *bufio.Reader) (interface{}, error){
	'+': func(re *bufio.Reader) (interface{}, error) {
		return readUntilDelimiter(re, delimiter)
	},
	'-': func(re *bufio.Reader) (interface{}, error) {
		return readUntilDelimiter(re, delimiter)
	},
	':': func(re *bufio.Reader) (interface{}, error) {
		return readInteger(re)
	},
	'$': func(re *bufio.Reader) (interface{}, error) {
		var buf []byte
		val, err := readInteger(re)
		if err == nil {
			buf = make([]byte, val)
			_, err = re.Read(buf)
		}
		return string(buf), err
	},
	// Not parsing array since it should be recursive! implement in handleConnection
}

type RespHandler struct {
}

func (rh *RespHandler) HandleConnection(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	data, err := readUntilDelimiter(reader, delimiter)

	if err != nil {
		log.Error(err)
		return err
	}

	command := rh.ParseMessage(data)
	return nil
}

func readUntilDelimiter(re *bufio.Reader, delimiter string) (string, error) {
	buf := make([]byte, 64, 0)
	delimiterLen := len(delimiter)

	for {
		msg, err := re.ReadString(delimiter[delimiterLen-1])

		if err != nil {
			log.Error(fmt.Sprintf("Error While Parsing: %v, %v", msg, err.Error()))
			return "", err
		}

		buf = append(buf, []byte(msg)...)
		if bytes.HasSuffix(buf, []byte(delimiter)) {
			return string(buf[:len(buf)-len(delimiter)]), nil
		}
	}
}

func readInteger(re *bufio.Reader) (int, error) {
	res := 0
	str, err := readUntilDelimiter(re, delimiter)
	if err == nil {
		res, err = strconv.Atoi(str)
	}
	return res, err
}

func (rh *RespHandler) ParseMessage(re *bufio.Reader) (msg string, err error) {
	valType, err := re.ReadByte()
	if err == nil {
		convertToVal[valType](re)
	}
}
