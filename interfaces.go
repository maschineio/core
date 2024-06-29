package core

type Replaceable interface {
	String() string        // string representation of the key / path
	Key() string           // key is the original key
	JSONPath() string      // returns the jsonpath
	IsContextObject() bool // returns true if the object is a context object
}
