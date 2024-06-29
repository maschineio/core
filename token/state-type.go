package token

// StateType represents all known states of the state machine.
type StateType int

const (
	IllegalStateType StateType = iota
	TaskType
	MapType
	ParallelType
	ChoiceType
	PassType
	WaitType
	SucceedType
	FailType
)
