package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	Rules         map[string]Rule
	MyTicket      []int
	NearbyTickets [][]int
}

func NewProgram() *Program {
	return &Program{
		Rules:         make(map[string]Rule),
		MyTicket:      make([]int, 0),
		NearbyTickets: make([][]int, 0),
	}
}

var (
	ruleMatcher *regexp.Regexp
)

// Rule represents the two ranges
// in which a number should be
type Rule struct {
	FirstRange  [2]int
	SecondRange [2]int
}

// NewRule creates a new rule
func NewRule(firstRange [2]int, secondRange [2]int) Rule {
	return Rule{
		FirstRange:  firstRange,
		SecondRange: secondRange,
	}
}

// Valid checks that the number is between those numbers
func (r Rule) Valid(n int) bool {
	return (r.FirstRange[0] <= n && n <= r.FirstRange[1]) || (r.SecondRange[0] <= n && n <= r.SecondRange[1])
}

func FirstPart(lines []string) (int, error) {
	p, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	errRate := 0
	for _, ticket := range p.NearbyTickets {
		for _, n := range ticket {
			var valid bool
			for _, r := range p.Rules {
				if r.Valid(n) {
					valid = true
				}
			}
			if !valid {
				errRate += n
				break
			}
		}
	}
	return errRate, nil
}

func remove(tickets [][]int, toRemove []int) [][]int {
	match := make(map[int]bool)
	for _, v := range toRemove {
		match[v] = true
	}
	validTickets := make([][]int, 0)
	for idx, t := range tickets {
		if _, ok := match[idx]; ok {
			continue
		}
		validTickets = append(validTickets, t)
	}
	return validTickets
}

func SecondPart(lines []string, keyword string) (int, error) {
	p, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	toRemove := []int{}
	for idx, ticket := range p.NearbyTickets {
		for _, n := range ticket {
			var valid bool
			for _, r := range p.Rules {
				if r.Valid(n) {
					valid = true
				}
			}
			if !valid {
				toRemove = append(toRemove, idx)
			}
		}
	}
	validTickets := remove(p.NearbyTickets, toRemove)
	fieldToPossibleIdx := map[string]map[int]bool{}
	for name, field := range p.Rules {
		idxs := map[int]bool{}
		for idx := range validTickets[0] {
			idxs[idx] = true
		}
		for _, ticket := range validTickets {
			for idx, i := range ticket {
				idxs[idx] = idxs[idx] && field.Valid(i)
			}
		}
		fieldToPossibleIdx[name] = idxs
	}

	fieldToIdx := map[string]int{}
	for {
		idx := -1
		for field, possibleIdx := range fieldToPossibleIdx {
			if truthyValues(possibleIdx) == 1 {
				idx = getTruthyValue(possibleIdx)
				fieldToIdx[field] = idx
			}
		}
		if idx < 0 {
			break

		}
		for _, possibleIdx := range fieldToPossibleIdx {
			possibleIdx[idx] = false
		}
	}

	result := 1
	for k := range p.Rules {
		if strings.HasPrefix(k, "departure ") {
			result *= p.MyTicket[fieldToIdx[k]]
		}
	}

	return result, nil
}

func truthyValues(m map[int]bool) int {
	r := 0
	for _, v := range m {
		if v {
			r++
		}
	}
	return r
}

func getTruthyValue(m map[int]bool) int {
	for idx, v := range m {
		if v {
			return idx
		}
	}
	return -1
}

func readLines(lines []string) (*Program, error) {
	p := NewProgram()
	phase := 0
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if l == "" {
			phase++
			i++
			continue
		}
		switch phase {
		case 0:
			err := addRule(p, l)
			if err != nil {
				return nil, err
			}
		case 1:
			ticket, err := parseTicket(l)
			if err != nil {
				return nil, err
			}
			p.MyTicket = ticket
		case 2:
			ticket, err := parseTicket(l)
			if err != nil {
				return nil, err
			}
			p.NearbyTickets = append(p.NearbyTickets, ticket)
		default:
			return nil, fmt.Errorf("not a valid phase: %d", phase)
		}
	}

	return p, nil
}

func parseTicket(l string) ([]int, error) {
	split := strings.Split(l, ",")
	result := make([]int, len(split))
	for i, s := range split {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}
	return result, nil
}

func addRule(p *Program, l string) error {
	matches := ruleMatcher.FindStringSubmatch(l)

	name := matches[1]
	firstRange1 := matches[2]
	firstRange2 := matches[3]
	secondRange1 := matches[4]
	secondRange2 := matches[5]

	var (
		firstRange  [2]int
		secondRange [2]int
		err         error
	)
	firstRange[0], err = strconv.Atoi(firstRange1)
	if err != nil {
		return err
	}
	firstRange[1], err = strconv.Atoi(firstRange2)
	if err != nil {
		return err
	}
	secondRange[0], err = strconv.Atoi(secondRange1)
	if err != nil {
		return err
	}
	secondRange[1], err = strconv.Atoi(secondRange2)
	if err != nil {
		return err
	}
	p.Rules[name] = NewRule(firstRange, secondRange)
	return nil
}

func init() {
	ruleMatcher = regexp.MustCompile(`([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)
}
