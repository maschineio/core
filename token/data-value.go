package token

import (
	"fmt"
	"strings"
)

// KV holds a key/value tuple
type KV struct {
	Key   string
	Value any
}

func (kv *KV) String() string {
	return fmt.Sprintf("Key(%s)=Val(%T, '%+v')", kv.Key, kv.Value, kv.Value)
}

// ReplacementKV contains ReplacementKey/Value of Type T
type ReplacementKV[T any] struct {
	Key   ReplacementKey
	Value DataValue[T]
}

// ReplacementKey
type ReplacementKey interface {
	Key() string       // returns the key without .$ at the end
	OriginKey() string // returns the original replacement key like myKey.$
}

type replacementKey struct {
	key    string
	origin string
}

func (r replacementKey) Key() string {
	return r.key
}

func (r replacementKey) OriginKey() string {
	return r.origin
}

func NewReplacementKey(origin string) ReplacementKey {
	key := origin
	if strings.HasSuffix(origin, ".$") {
		key = origin[:len(origin)-2]
	}
	return replacementKey{key: key, origin: origin}
}

type DataValue[T any] interface {
	Type() DataType
	Value() T
}

type dataValue[T any] struct {
	value    T
	dataType DataType
}

func (dv dataValue[T]) Type() DataType {
	return dv.dataType
}

func (dv dataValue[T]) Value() T {
	return dv.value
}

func NewDataValue[T any](value T, t DataType) DataValue[T] {
	return dataValue[T]{value: value, dataType: t}
}
