package DB

import (
	"errors"
	"fmt"
	"sync"
)

type VolatileDB struct {
	sync.Mutex
	data map[interface{}]interface{}
}

func (v *VolatileDB) Get(key interface{}) (interface{}, error) {
	v.Lock()
	defer v.Unlock()

	if val, ok := v.data[key]; ok {
		return val, nil
	}

	return nil, errors.New(fmt.Sprintf("Unkown key: %v", key))
}

func (v *VolatileDB) Set(key interface{}, val interface{}) (succeeded bool) {
	v.Lock()
	defer v.Unlock()

	v.data[key] = val
	return true
}
