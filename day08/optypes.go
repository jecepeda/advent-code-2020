package day08

import "fmt"

// OpType represents the operations we can get on the
// input lines
type OpType int

// values OpType can get
const (
	Acc  OpType = iota // increment
	Jmp                // jump to line
	Noop               // no operation, go to next line
)

var opTypes = [...]string{
	Acc:  "acc",
	Jmp:  "jmp",
	Noop: "nop",
}

// MarshalText renders the message type as text
func (mt OpType) MarshalText() (text []byte, err error) {
	return []byte(mt.String()), nil
}

// UnmarshalText parses a message type from a textual representation
func (mt *OpType) UnmarshalText(text []byte) error {
	for i, ac := range opTypes {
		if ac == string(text) {
			*mt = OpType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown message type: %v", string(text))
}

func (mt OpType) String() string {
	return opTypes[mt]
}
