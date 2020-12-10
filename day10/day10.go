package day10

import "sort"

func sortAndAddAdaptors(lines []int) []int {
	lines = append(lines, 0)
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	lines = append(lines, lines[len(lines)-1]+3)
	return lines
}

func FirstPart(lines []int) int {
	var (
		ones   int
		threes int
	)
	lines = sortAndAddAdaptors(lines)
	for i := 1; i < len(lines); i++ {
		if diff := lines[i] - lines[i-1]; diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
	}
	return ones * threes
}

func SecondPart(lines []int) int {
	lines = sortAndAddAdaptors(lines)
	m := make(map[int]bool)
	for i := range lines {
		m[lines[i]] = true
	}
	return findCombinations(lines, m, make(map[int]int), lines[0])
}

func findCombinations(lines []int, lineElemMap map[int]bool, partialResults map[int]int, n int) int {
	if lines[len(lines)-1] == n {
		return 1
	}
	if !lineElemMap[n] {
		return 0
	}
	res, ok := partialResults[n]
	if !ok {
		res = findCombinations(lines, lineElemMap, partialResults, n+1) + findCombinations(lines, lineElemMap, partialResults, n+2) + findCombinations(lines, lineElemMap, partialResults, n+3)
		partialResults[n] = res
	}
	return res
}
