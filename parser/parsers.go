package parser

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Parser interface {
	ParseMessage(re *bufio.Reader) (result interface{}, err error)
}

func readInteger(re *bufio.Reader) (int, error) {
	res := 0
	str, err := readUntilDelimiter(re, delimiter)
	if err == nil {
		res, err = strconv.Atoi(str)
	}
	return res, err
}

func readUntilDelimiter(re *bufio.Reader, delimiter string) (string, error) {
	buf := make([]byte, 64, 0)
	delimiterLen := len(delimiter)

	for {
		msg, err := re.ReadString(delimiter[delimiterLen-1])

		if err != nil {
			log.Error(fmt.Sprintf("Error While Reading: %v, %v", msg, err.Error()))
			return "", err
		}

		buf = append(buf, []byte(msg)...)
		if bytes.HasSuffix(buf, []byte(delimiter)) {
			return string(buf[:len(buf)-len(delimiter)]), nil
		}
	}
}
