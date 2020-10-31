package main

import (
	"bufio"
	"goRedis/parser"
	"goRedis/server"
)

func main() {

	pr := new(parser.RESPParser)
	pr.Parsers = map[byte]func(re *bufio.Reader) (interface{}, error){
		'+': pr.StringParser,
		'-': pr.StringParser,
		':': pr.IntParser,
		'$': pr.BulkStringParser,
		'*': pr.ArrayParser,
	}

	var se server.Server = server.NewTCPServer(5535, pr)
	se.StartListen()
}
