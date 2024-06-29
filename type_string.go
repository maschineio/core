// Code generated by "stringer -type=Type pkg/core/type.go"; DO NOT EDIT.

package core

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Nil-1]
	_ = x[String-2]
	_ = x[Bool-3]
	_ = x[Bytes-4]
	_ = x[Float-5]
	_ = x[StringMap-6]
	_ = x[PointerStringMap-7]
	_ = x[Slice-8]
}

const _Type_name = "UnknownNilStringBoolBytesFloatStringMapPointerStringMapSlice"

var _Type_index = [...]uint8{0, 7, 10, 16, 20, 25, 30, 39, 55, 60}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
