package redis

import "goRedis/redis/DB"

type Redis struct {
	Commands map[interface{}]func(params []interface{}) (interface{}, error)
	DB DB.DB
}
