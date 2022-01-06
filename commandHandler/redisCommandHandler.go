package commandHandler

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"goRedis/redis"
)

type RedisCommandHandler struct {
	redis *redis.Redis
}

func NewRedisCommandHandler(redis *redis.Redis) RedisCommandHandler {
	return RedisCommandHandler{redis}
}

func (r RedisCommandHandler) HandleCommand(command interface{}) (interface{}, error) {
	arr, ok := command.([]interface{})
	if !ok || len(arr) == 0 {
		msg := fmt.Sprintf("%v is not an array which mean it not a command!", arr)
		log.Error(msg)
		return nil, errors.New(msg)
	}

	return r.redis.Commands[arr[0]](arr[1:])
}
