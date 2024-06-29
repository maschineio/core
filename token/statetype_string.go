// Code generated by "stringer -type=StateType token/state-type.go"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IllegalStateType-0]
	_ = x[TaskType-1]
	_ = x[MapType-2]
	_ = x[ParallelType-3]
	_ = x[ChoiceType-4]
	_ = x[PassType-5]
	_ = x[WaitType-6]
	_ = x[SucceedType-7]
	_ = x[FailType-8]
}

const _StateType_name = "IllegalStateTypeTaskTypeMapTypeParallelTypeChoiceTypePassTypeWaitTypeSucceedTypeFailType"

var _StateType_index = [...]uint8{0, 16, 24, 31, 43, 53, 61, 69, 80, 88}

func (i StateType) String() string {
	if i < 0 || i >= StateType(len(_StateType_index)-1) {
		return "StateType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StateType_name[_StateType_index[i]:_StateType_index[i+1]]
}
