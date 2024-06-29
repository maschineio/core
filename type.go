package core

type Type int

const (
	Unknown Type = iota
	Nil
	String
	Bool
	Bytes
	Float
	StringMap
	PointerStringMap
	Slice
)
