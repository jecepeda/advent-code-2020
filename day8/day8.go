package day8

import (
	"errors"
	"strconv"
	"strings"
)

// ProgramStatus represents the status of the program
type ProgramStatus int

// values ProgramStatus can get
const (
	NotRunning ProgramStatus = iota
	Halted
	Success
)

// Program represents a set of operations
// and the final status of the program
type Program struct {
	Operations []Operation
	Status     ProgramStatus
}

// NewProgram creates a new program
func NewProgram(operations []Operation) Program {
	ops := make([]Operation, len(operations))
	copy(ops, operations)
	return Program{
		Operations: ops,
	}
}

// Run runs the program, returning the last value of the accumulator
// and the position of the program on its last execution
func (o *Program) Run() (acc int, pos int) {
	var (
		incAcc, incPos int
	)
	for {
		if pos >= len(o.Operations) {
			o.Status = Success
			return
		} else if o.Operations[pos].HasBeenCalled() {
			o.Status = Halted
			return
		}
		incAcc, incPos = o.Operations[pos].Exec()
		acc += incAcc
		pos += incPos
	}
}

// Operation represents an operation
type Operation struct {
	Type        OpType
	Value       int
	TimesCalled int
}

// NewOperation initializes a new operation
func NewOperation(opType OpType, value int) Operation {
	return Operation{
		Type:  opType,
		Value: value,
	}
}

// HasBeenCalled checks if the operation has been called
// at least once
func (o *Operation) HasBeenCalled() bool {
	return o.TimesCalled > 0
}

// Exec executes the operation, returning the increment into
// the accumulator, and the line position in which
// the program has to jump
func (o *Operation) Exec() (accIncrement, lineIncrement int) {
	o.TimesCalled++
	switch o.Type {
	case Acc:
		return o.Value, 1
	case Jmp:
		return 0, o.Value
	case Noop:
		return 0, 1
	}
	return 0, 0
}

// FirstPart returns the value of the accumulator when one operation has been executed at lest once
func FirstPart(lines []string) (int, error) {
	operations, err := readProgramLines(lines)
	if err != nil {
		return 0, err
	}
	p := NewProgram(operations)
	acc, _ := p.Run()
	return acc, nil
}

// SecondPart attempts to fix the program, switching noop and jmp values
func SecondPart(lines []string) (int, error) {
	operations, err := readProgramLines(lines)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(operations); i++ {
		if operations[i].Type == Acc {
			continue
		}
		operations[i] = switchOperation(operations[i])
		p := NewProgram(operations)
		acc, _ := p.Run()
		if p.Status == Success {
			return acc, nil
		}
		// go back to normal
		operations[i] = switchOperation(operations[i])
	}
	return 0, errors.New("could not fix the program")
}

func switchOperation(o Operation) Operation {
	if o.Type == Noop {
		o.Type = Jmp
		return o
	}
	o.Type = Noop
	return o
}

func readProgramLines(lines []string) ([]Operation, error) {
	var (
		op         OpType
		value      int
		err        error
		operations []Operation
	)
	for _, l := range lines {
		parts := strings.Split(l, " ")
		err = op.UnmarshalText([]byte(parts[0]))
		if err != nil {
			return nil, err
		}
		value, err = strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		operations = append(operations, NewOperation(op, value))
	}
	return operations, nil
}
