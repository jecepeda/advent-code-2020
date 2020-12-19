package day18

import (
	"regexp"
	"strconv"
)

var (
	tokenMatcher *regexp.Regexp
)

// Kind represents
type Kind int

// values Kind can get
const (
	Number Kind = iota
	Sum
	Mul
	RightParen
	LeftParen
)

// Token represents the capacity
// and the use of its content
// it allows to distinguis between operations (*,+)
// integer values, and parenthesis
type Token struct {
	Kind   Kind
	Number int
}

func readLines(lines []string) ([][]Token, error) {
	result := [][]Token{}
	for i := range lines {
		matches := tokenMatcher.FindAllString(lines[i], -1)
		tokens := make([]Token, 0)
		for _, m := range matches {
			switch m {
			case "(":
				tokens = append(tokens, Token{Kind: LeftParen})
			case ")":
				tokens = append(tokens, Token{Kind: RightParen})
			case "*":
				tokens = append(tokens, Token{Kind: Mul})
			case "+":
				tokens = append(tokens, Token{Kind: Sum})
			default:
				v, err := strconv.Atoi(m)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, Token{Kind: Number, Number: v})
			}
		}
		result = append(result, tokens)
	}
	return result, nil
}

func generateRPNQueue(tokens []Token) Queue {
	precedences := map[Kind]int{
		Mul:    1,
		Sum:    2,
		Number: 10,
	}
	s := Stack{}
	output := Queue{}
	for _, t := range tokens {
		switch t.Kind {
		case LeftParen:
			s.Push(t)
		case RightParen:
			for s.Peek().Kind != LeftParen {
				output.Add(s.Pop())
			}
			s.Pop()
		case Mul, Sum:
			priority := precedences[t.Kind]
			for s.Len() > 0 && priority <= precedences[s.Peek().Kind] {
				output.Add(s.Pop())
			}
			s.Push(t)
		case Number:
			s.Push(t)
		}
	}
	for s.Len() > 0 {
		output.Add(s.Pop())
	}
	return output
}

func computeRPN(q *Queue) int {
	s := Stack{}
	for q.Len() > 0 {
		v := q.Dequeue()
		switch v.Kind {
		case Number:
			s.Push(v)
		case Mul:
			v1, v2 := s.Pop(), s.Pop()
			r := v1.Number * v2.Number
			s.Push(Token{Kind: Number, Number: r})
		case Sum:
			v1, v2 := s.Pop(), s.Pop()
			r := v1.Number + v2.Number
			s.Push(Token{Kind: Number, Number: r})
		}
	}
	return s.Pop().Number
}

func compute(q *Queue) int {
	s := Stack{}
	for q.Len() > 0 {
		v := q.Dequeue()
		switch v.Kind {
		case Number:
			s.Push(v)
		case Sum:
			left := s.Pop()
			right := q.Dequeue()
			var rValue int
			switch right.Kind {
			case Number:
				rValue = right.Number
			case LeftParen:
				rValue = compute(q)
			}
			v := left.Number + rValue
			s.Push(Token{Kind: Number, Number: v})
		case Mul:
			left := s.Pop()
			right := q.Dequeue()
			var rValue int
			switch right.Kind {
			case Number:
				rValue = right.Number
			case LeftParen:
				rValue = compute(q)
			}
			v := left.Number * rValue
			s.Push(Token{Kind: Number, Number: v})
		case LeftParen:
			v := compute(q)
			s.Push(Token{Kind: Number, Number: v})
		case RightParen:
			return s.Pop().Number
		}
	}
	return s.Pop().Number
}

func FirstPart(lines []string) (int, error) {
	calculations, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, c := range calculations {
		q := Queue{Data: c}
		result += compute(&q)
	}
	return result, nil
}

func SecondPart(lines []string) (int, error) {
	calculations, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, c := range calculations {
		q := generateRPNQueue(c)
		result += computeRPN(&q)
	}
	return result, nil
}

func init() {
	tokenMatcher = regexp.MustCompile(`(\d+|\*|\+|\(|\))`)
}
