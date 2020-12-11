package day01

// Sum2020Part1 finds those two numbers whose sum is 2020.
// returns the multiplication of those two numbers
func Sum2020Part1(numbers []int) int {
	difference := make(map[int]int)
	for _, n := range numbers {
		difference[2020-n] = n
	}
	for _, n := range numbers {
		v, ok := difference[n]
		if ok {
			return v * n
		}
	}
	return 0
}

// Sum2020Part2 finds those three numbers whose sum is 2020.
// returns the multiplication of those three numbers
func Sum2020Part2(numbers []int) int {
	difference := make(map[int]int)
	for _, n := range numbers {
		difference[2020-n] = n
	}
	var (
		i, first, second, third int
		ok                      bool
	)
	for i, second = range numbers {
		for j := i; j < len(numbers); j++ {
			third = numbers[j]
			first, ok = difference[second+third]
			if ok {
				return first * second * third
			}
		}
	}
	return 0
}
