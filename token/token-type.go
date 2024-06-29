package token

type TokenType int

const (
	// comparators
	StringEquals TokenType = (iota << 1) + 2
	StringLessThan
	StringGreaterThan
	StringLessThanEquals
	StringGreaterThanEquals
	StringMatches
	NumericEquals
	NumericLessThan
	NumericGreaterThan
	NumericLessThanEquals
	NumericGreaterThanEquals
	BooleanEquals
	TimestampEquals
	TimestampLessThan
	TimestampGreaterThan
	TimestampLessThanEquals
	TimestampGreaterThanEquals
	IsNull
	IsPresent
	IsNumeric
	IsString
	IsBoolean
	IsTimestamp
	StringEqualsPath
	StringLessThanPath
	StringGreaterThanPath
	StringLessThanEqualsPath
	StringGreaterThanEqualsPath
	NumericEqualsPath
	NumericLessThanPath
	NumericGreaterThanPath
	NumericLessThanEqualsPath
	NumericGreaterThanEqualsPath
	BooleanEqualsPath
	TimestampEqualsPath
	TimestampLessThanPath
	TimestampGreaterThanPath
	TimestampLessThanEqualsPath
	TimestampGreaterThanEqualsPath // end of standard comparators
	StringLenEquals                // extended comparators
	StringLenLessThan
	StringLenGreaterThan
	StringLenLessThanEquals
	StringLenGreaterThanEquals
	IsStringEmpty
	ArrayLenEquals
	ArrayLenLessThan
	ArrayLenGreaterThan
	ArrayLenLessThanEquals
	ArrayLenGreaterThanEquals
	IsArrayEmpty
	IsArray
	StringLenEqualsPath
	StringLenLessThanPath
	StringLenGreaterThanPath
	StringLenLessThanEqualsPath
	StringLenGreaterThanEqualsPath
	ArrayLenEqualsPath
	ArrayLenLessThenPath
	ArrayLenGreaterThanPath
	ArrayLenLessThanEqualsPath
	ArrayLenGreaterThanEqualsPath // last extended comparator
	_                             // make room for extensions
	_
	_
	ILLEGAL
	Comment
	End
	False
	Next
	Null
	Resource
	StartAt
	TimeoutSeconds
	TimeoutSecondsPath
	True
	Version
	InputPath
	OutputPath
	ResultPath
	ResultSelector
	Type
	States
	Id
	Parameters
	HeartbeatSeconds
	HeartbeatSecondsPath
	Result
	Credentials
	Seconds
	SecondsPath
	Timestamp
	TimestampPath
	Error
	ErrorPath
	Cause
	CausePath
	ErrorEquals
	IntervalSeconds
	MaxAttempts
	BackoffRate
	MaxDelaySeconds
	JitterStrategy
	Retry
	Catch
	Choices
	Default
	Variable
	SimpleBoolRule
	AndBoolRule
	OrBoolRule
	NotBoolRule
	Branches
	ToleratedFailureCountPath
	ToleratedFailurePercentagePath
	ToleratedFailureCount
	ToleratedFailurePercentage
	ItemSelector
	ItemBatcher
	MaxConcurrency
	MaxConcurrencyPath
	ProcessorConfig
	ItemProcessor
	ItemsPath
	ContextObject
	StatesFunction
	MathFunction
	ResultWriter
	UnquotedJSONPath
	HashFunction
	SingleQuotedString
	SingleQuotedParameters
	TemplateParameter
	MaxItems
	MaxItemsPath
	InputType
	CSVHeaderLocation
	ItemReader
	ReaderConfig
	BatchInput
	MaxItemsPerBatch
	MaxItemsPerBatchPath
	MaxInputBytesPerBatch
	MaxInputBytesPerBatchPath
)
