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

func (v *VolatileDB) Get(key string) interface{} {
	v.Lock()
	defer v.Unlock()

	if val, ok := v.data[key]; ok {
		return val
	}

	return errors.New(fmt.Sprintf("Unkown key: %v", key))
}

func (v *VolatileDB) Set(key string, val interface{}) (succeeded bool) {
	v.Lock()
	defer v.Unlock()

	v.data[key] = val
	return true
}
