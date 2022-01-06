package DB

type DB interface {
	Get(key interface{}) (interface{}, error)
	Set(key interface{}, val interface{}) (succeeded bool)
}
