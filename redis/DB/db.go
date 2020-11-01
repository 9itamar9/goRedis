package DB

type DB interface {
	Get(key string) interface{}
	Set(key string, val interface{}) (succeeded bool)
}
