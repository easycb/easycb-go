package easycb

import (
	"encoding/json"
	"sync"
)

type AnyMap map[string]interface{}

var mu = new(sync.RWMutex)

func (am AnyMap) Set(key string, value interface{}) AnyMap {
	mu.Lock()
	am[key] = value
	mu.Unlock()
	return am
}

func (am AnyMap) GetString(key string) string {
	if am == nil {
		return ""
	}
	mu.RLock()
	defer mu.RUnlock()
	value, ok := am[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

func (am AnyMap) GetInterface(key string) interface{} {
	if am == nil {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()
	return am[key]
}

func (am AnyMap) Remove(key string) {
	mu.Lock()
	delete(am, key)
	mu.Unlock()
}

func (am AnyMap) Reset() {
	mu.Lock()
	for k := range am {
		delete(am, k)
	}
	mu.Unlock()
}

func (am AnyMap) JsonBody() (jb string) {
	mu.Lock()
	defer mu.Unlock()
	bs, err := json.Marshal(am)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
