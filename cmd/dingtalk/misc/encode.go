package misc

import (
	"encoding/json"
)

type Params map[string][]interface{}

func (p Params) Set(key string, value interface{}) {
	p[key] = []interface{}{value}

}

func (p Params) Add(key string, value interface{}) {
	p[key] = append(p[key], value)
}

func (p Params) Del(key string) {
	delete(p, key)
}

func (p Params) Get(key string) interface{} {
	if p == nil {
		return nil
	}
	if pi := p[key]; len(pi) == 0 {
		return nil
	} else {
		return pi[0]
	}
}

// convert map to json
func (p Params) EncodeToJson() ([]byte, error) {
	if res, err := json.Marshal(p); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
