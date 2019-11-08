package misc

import (
	"encoding/json"
)

type Data map[string]interface{}

func (d Data) Set(key string, value interface{}) {
	switch value.(type) {
	case []interface{}:
		d[key] = value.([]interface{})
		return
	}
	d[key] = value
}

func (d Data) Add(key string, value interface{}) {
}

func (d Data) Del(key string) {
	delete(d, key)
}

func (d Data) Get(key string) interface{} {
	if d == nil {
		return nil
	}
	switch d[key].(type) {
	case []interface{}:
		if len(d[key].([]interface{})) != 0 {
			return d[key].([]interface{})[0]
		}
	}
	if d[key] == nil {
		return nil
	} else {
		return d[key]
	}
}

// convert map to json
func (d Data) EncodeToJson() ([]byte, error) {
	if res, err := json.Marshal(d); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (d Data) EncodeToURLParameter() string {
	return ""
}
