package core

import (
	"fmt"
)

type contextObject struct {
	key             string
	jsonPath        string
	isContextObject bool
}

func (co contextObject) String() string {
	return fmt.Sprintf("context(%s).path(%s)", co.key, co.jsonPath)
}

func (co contextObject) Key() string {
	return co.key
}

func (co contextObject) JSONPath() string {
	return co.jsonPath
}

func (co contextObject) IsContextObject() bool {
	return co.isContextObject
}

func NewContextObject(key string) Replaceable {
	return contextObject{key: key, jsonPath: key[1:], isContextObject: true}
}

func NewJSONPath(key string) Replaceable {
	return contextObject{key: key, jsonPath: key, isContextObject: false}
}
