package day14

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Program struct {
	Mask   string
	Memory map[int]int
}

// NewProgram creates a new program
func NewProgram() *Program {
	return &Program{
		Mask:   "0",
		Memory: make(map[int]int),
	}
}

// SetMask sets a new mask, to filter incoming values
func (p *Program) SetMask(mask string) {
	p.Mask = mask
}

// AddToMemory adds a value into a memory position, applying a bit mask
func (p *Program) AddToMemory(memPos int, value int) error {
	mFormat := "%0" + fmt.Sprint(len(p.Mask)) + "b"
	binaryRepresentation := []rune(fmt.Sprintf(mFormat, value))
	for i, v := range p.Mask {
		if v != 'X' {
			binaryRepresentation[i] = v
		}
	}
	v, err := strconv.ParseInt(string(binaryRepresentation), 2, 64)
	if err != nil {
		return err
	}
	p.Memory[memPos] = int(v)

	return err
}

// DecodeAddresses decodes the memory code using the bitmask, getting
// all the necessary addresses, and setting the value to those addresses
func (p *Program) DecodeAddresses(memCode, value int) error {
	mFormat := "%0" + fmt.Sprint(len(p.Mask)) + "b"
	binaryRepresentation := []rune(fmt.Sprintf(mFormat, memCode))
	for i, v := range p.Mask {
		if v == 'X' {
			binaryRepresentation[i] = 'X'
		} else if v == '1' {
			binaryRepresentation[i] = '1'
		}
	}
	addresses, err := p.getMemoryAddresses(binaryRepresentation)
	if err != nil {
		return err
	}
	for _, a := range addresses {
		p.Memory[a] = value
	}

	return nil
}

func (p *Program) getMemoryAddresses(code []rune) ([]int, error) {
	xPositions := make([]int, 0)
	for i, v := range code {
		if v == 'X' {
			xPositions = append(xPositions, i)
		}
	}
	numPosibilities := int(math.Pow(2, float64(len(xPositions))))
	replacement := make([]bool, len(xPositions))
	combinations := make([][]rune, numPosibilities)
	for i := 0; i < numPosibilities; i++ {
		getPossibility(replacement, i)
		combinations[i] = make([]rune, len(code))
		copy(combinations[i], code)
		for xPos, v := range replacement {
			if v {
				combinations[i][xPositions[xPos]] = '1'
			} else {
				combinations[i][xPositions[xPos]] = '0'
			}
		}
	}
	result := make([]int, len(combinations))
	for i := range combinations {
		v, err := strconv.ParseInt(string(combinations[i]), 2, 64)
		if err != nil {
			return nil, err
		}
		result[i] = int(v)
	}
	return result, nil
}

func getPossibility(replacements []bool, n int) {
	for i := range replacements {
		if n&(1<<i) != 0 {
			replacements[i] = true
		} else {
			replacements[i] = false
		}
	}
}

// GetValueSum sums the memory values and return the result
func (p *Program) GetValueSum() int {
	result := 0
	for _, v := range p.Memory {
		result += v
	}
	return result
}

// FirstPart runs the program an returns the sum of all memorya ddresses
func FirstPart(lines []string) (int, error) {
	maskRegex, err := regexp.Compile(`mask = ([\dX]+)`)
	if err != nil {
		return 0, err
	}
	memRegex, err := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
	if err != nil {
		return 0, err
	}
	p := NewProgram()
	for _, l := range lines {
		if maskRegex.MatchString(l) {
			matches := maskRegex.FindStringSubmatch(l)
			p.SetMask(matches[1])
		} else {
			matches := memRegex.FindStringSubmatch(l)
			memPos, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, err
			}
			memValue, err := strconv.Atoi(matches[2])
			if err != nil {
				return 0, err
			}
			err = p.AddToMemory(memPos, memValue)
			if err != nil {
				return 0, err
			}
		}
	}
	return p.GetValueSum(), nil
}

// SecondPart runs the program an returns the sum of all memorya ddresses
func SecondPart(lines []string) (int, error) {
	maskRegex, err := regexp.Compile(`mask = ([\dX]+)`)
	if err != nil {
		return 0, err
	}
	memRegex, err := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
	if err != nil {
		return 0, err
	}
	p := NewProgram()
	for _, l := range lines {
		if maskRegex.MatchString(l) {
			matches := maskRegex.FindStringSubmatch(l)
			p.SetMask(matches[1])
		} else {
			matches := memRegex.FindStringSubmatch(l)
			memPos, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, err
			}
			memValue, err := strconv.Atoi(matches[2])
			if err != nil {
				return 0, err
			}
			err = p.DecodeAddresses(memPos, memValue)
			if err != nil {
				return 0, err
			}
		}
	}
	return p.GetValueSum(), nil
}
