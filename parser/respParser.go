package parser

import "bufio"

const delimiter = "\\r\\n"

type RESPParser struct {
	Parsers map[byte]func(re *bufio.Reader) (interface{}, error)
}

// NOTE: This parser still cannot handle invalid structure and will prob stuck waiting for the delimiter
func (pr *RESPParser) ParseMessage(re *bufio.Reader) (result interface{}, err error) {
	valType, err := re.ReadByte()
	if err == nil {
		result, err = pr.Parsers[valType](re)
	}

	return result, err
}

func (pr *RESPParser) StringParser(re *bufio.Reader) (interface{}, error) {
	return readUntilDelimiter(re, delimiter)
}

func (pr *RESPParser) BulkStringParser(re *bufio.Reader) (interface{}, error) {
	var buf []byte
	length, err := readInteger(re)
	if err == nil {
		buf = make([]byte, length + len(delimiter))
		_, err = re.Read(buf)
	}
	return string(buf[:length]), err
}

func (pr *RESPParser) IntParser(re *bufio.Reader) (interface{}, error) {
	return readInteger(re)
}

func (pr *RESPParser) ArrayParser(re *bufio.Reader) (interface{}, error) {
	arrLen, err := readInteger(re)
	if err != nil {
		return nil, err
	}

	arr := make([]interface{}, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i], err = pr.ParseMessage(re)
		if err != nil {
			return nil, err
		}
	}

	return arr, nil
}
