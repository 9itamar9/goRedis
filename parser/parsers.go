package parser

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const delimiter = "\r\n"

var convertToVal = map[byte]func(re *bufio.Reader) (interface{}, error){
	'+': StringParser,
	'-': StringParser,
	':': IntParser,
	'$': BulkStringParser,
	'*': ArrayParser,
}

func ParseMessage(re *bufio.Reader) (result interface{}, err error) {
	valType, err := re.ReadByte()
	if err == nil {
		result, err = convertToVal[valType](re)
	}

	return result, err
}

func StringParser(re *bufio.Reader) (interface{}, error) {
	return readUntilDelimiter(re, delimiter)
}

func BulkStringParser(re *bufio.Reader) (interface{}, error) {
	var buf []byte
	length, err := readInteger(re)
	if err == nil {
		buf = make([]byte, length)
		_, err = re.Read(buf)
	}
	return string(buf), err
}

func IntParser(re *bufio.Reader) (interface{}, error) {
	return readInteger(re)
}

func ArrayParser(re *bufio.Reader) (interface{}, error) {
	arrLen, err := readInteger(re)
	if err != nil {
		return nil, err
	}

	arr := make([]interface{}, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i], err = ParseMessage(re)
		if err != nil {
			return nil, err
		}
	}

	return arr, nil
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
