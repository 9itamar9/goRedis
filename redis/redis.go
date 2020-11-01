package redis

type redis struct {
	commands map[string]func(command []interface{}) interface{}

}
