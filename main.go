package main

import (
	"bufio"
	"errors"
	"fmt"
	"goRedis/commandHandler"
	"goRedis/parser"
	"goRedis/redis"
	"goRedis/redis/DB"
	"goRedis/server"
	"sync"
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

	red := new(redis.Redis)
	red.Commands = map[interface{}]func(params []interface{}) (interface{}, error){
		"GET": func(params []interface{}) (interface{}, error) {
			if len(params) == 0 {
				return nil, errors.New("not enough arguments")
			}
			return red.DB.Get(params[0])
		},
		"SET": func(params []interface{}) (interface{}, error) {
			if len(params) != 2 {
				return nil, errors.New("not enough arguments")
			}
			if red.DB.Set(params[0], params[1]) {
				return "OK", nil
			}
			return "", errors.New(fmt.Sprintf("Could not set value %v for %v", params[1], params[0]))
		},
	}
	red.DB = &DB.VolatileDB{sync.Mutex{}, make(map[interface{}]interface{})}

	redisCmd := commandHandler.NewRedisCommandHandler(red)

	var se server.Server = server.NewTCPServer(5535, pr, redisCmd)
	se.StartListen()
}
