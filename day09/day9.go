package day09

import "sort"

func SecondPart(numbers []int, preamble int) int {
	num := FindNotMatchingNumber(numbers, preamble)

	sumNumbers := FindSumNumbers(numbers, num)
	sort.Slice(sumNumbers, func(i, j int) bool {
		return sumNumbers[i] > sumNumbers[j]
	})
	return sumNumbers[0] + sumNumbers[len(sumNumbers)-1]
}

func FindSumNumbers(numbers []int, num int) []int {
	for i := range numbers {
		var n = numbers[i]
		for j := i + 1; j < len(numbers) && n < num; j++ {
			n += numbers[j]
			if n == num {
				return numbers[i : j+1]
			}
		}
	}
	return nil
}

// FindNotMatchingNumber finds the number which is not possible
// to sum with the n preamble numbers
func FindNotMatchingNumber(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		if !matchNumber(numbers[i], numbers[i-preamble:i]) {
			return numbers[i]
		}
	}
	return -1
}

func matchNumber(n int, preamble []int) bool {
	for i := range preamble {
		for j := range preamble {
			if preamble[i]+preamble[j] == n {
				return true
			}
		}
	}
	return false
}
