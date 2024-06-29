package token

// DataType known by the parser.
type DataType int

const (
	String             DataType = iota // 0
	Any                                // 1
	Bool                               // 2
	ComparisonFunction                 // 3 currently unused
	ContextPath                        // 4
	Float64                            // 5
	Int                                // 6
	JSONPath                           // 7
	Slice                              // 8
	SliceTokenMap                      // 9
	State                              // 10
	StringMap                          // 11
	StringSlice                        // 12
	TimestampType                      // 13
	TokenMap                           // 14
	UInt64                             // 15
	Function                           // 16 - a Function represents a StatesFunction or a MathFunction
)
