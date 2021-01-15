package day23

import (
	"strconv"
	"strings"
)

type Circle struct {
	variables []int
}

func (c *Circle) PickFrom(pos int) []int {
	result := []int{}
	begin := pos + 1
	end := pos + 4

	if end > len(c.variables) {
		circleSize := len(c.variables)
		difference := end - circleSize
		result = append(result, c.variables[begin:circleSize]...)
		result = append(result, c.variables[0:difference]...)
		c.variables = c.variables[difference:begin]
	} else {
		result = append(result, c.variables[begin:end]...)
		c.variables = join(c.variables[:begin], c.variables[end:])
	}
	return result
}

func (c *Circle) Insert(value, valuePos int, values []int) {
	var (
		left  []int
		right []int
	)
	for i, v := range c.variables {
		if v == value {
			left = append(left, c.variables[0:i]...)
			right = append(right, c.variables[i+1:]...)
			break
		}
	}
	found, whereToInsert, isLeft := getInsertPoint(value, left, right)
	if !found {
		_, whereToInsert, isLeft = getInsertPoint(max(join(left, right))+1, left, right)
	}
	if isLeft {
		left = append(left[0:whereToInsert+1], join(values, left[whereToInsert+1:])...)
	} else {
		right = append(right[0:whereToInsert+1], join(values, right[whereToInsert+1:])...)
	}
	// join everything
	final := append(left, join([]int{value}, right)...)
	for i, v := range final {
		if v == value && i != valuePos {
			diff := valuePos - i
			if diff < 0 {
				rest := final[0:-diff]
				final = final[-diff:]
				final = append(final, rest...)
			} else {
				rest := final[len(final)-diff:]
				final = final[:len(final)-diff]
				final = append(rest, final...)
			}
			break
		}
	}
	c.variables = final
}

func getInsertPoint(value int, left, right []int) (found bool, whereToInsert int, isLeft bool) {
	for i := value - 1; i > 0 && !found; i-- {
		for j := range left {
			if left[j] == i {
				found, isLeft = true, true
				whereToInsert = j
				break
			}
		}
		for j := range right {
			if right[j] == i {
				found = true
				whereToInsert = j
				break
			}
		}
	}
	return
}

func join(a, b []int) []int {
	return append(a, b...)
}

func max(l []int) int {
	var max = -1
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

func FirstPart(values []int, rounds int) string {
	valueLength := len(values)
	c := &Circle{
		variables: values,
	}
	for i := 0; i < rounds; i++ {
		vPos := i % valueLength
		vVal := c.variables[vPos]
		extracted := c.PickFrom(vPos)
		c.Insert(vVal, vPos, extracted)
	}
	var onePos int
	for i := range c.variables {
		if c.variables[i] == 1 {
			onePos = i
			break
		}
	}
	vars := append(c.variables[onePos+1:], c.variables[:onePos]...)
	var result strings.Builder
	for _, v := range vars {
		result.WriteString(strconv.Itoa(v))
	}
	return result.String()
}

func SecondPart(initialValues []int, size, rounds int) int {
	cups, first, max := buildCaps(initialValues, size)
	cups = playCups(cups, first, max, rounds)
	v1, v2 := cups[1].Next, cups[1].Next.Next
	return v1.Value * v2.Value
}

type Cup struct {
	Value int
	Next  *Cup
}

func playCups(cups map[int]*Cup, first *Cup, max, rounds int) map[int]*Cup {
	curr := first
	for t := 0; t < rounds; t++ {
		c1, c2, c3 := curr.Next, curr.Next.Next, curr.Next.Next.Next
		curr.Next = c3.Next

		n := curr.Value
		for {
			n--
			if n == 0 {
				n = max
			}
			if c1.Value != n && c2.Value != n && c3.Value != n {
				break
			}
		}
		dest := cups[n]
		c3.Next = dest.Next
		dest.Next = c1
		curr = curr.Next
	}
	return cups
}

func buildCaps(initialValues []int, finalLength int) (map[int]*Cup, *Cup, int) {
	var (
		first *Cup
		curr  *Cup
		cups  = make(map[int]*Cup)
		max   = 0
	)
	for _, v := range initialValues {
		if v > max {
			max = v
		}
		c := &Cup{
			Value: v,
		}
		if first == nil {
			first = c
			curr = c
		} else {
			curr.Next = c
			curr = c
		}
		cups[v] = c
	}
	for len(cups) < finalLength {
		max++
		c := &Cup{
			Value: max,
		}
		curr.Next = c
		curr = c
		cups[max] = c
	}
	curr.Next = cups[initialValues[0]] // close the circle
	return cups, first, max
}
