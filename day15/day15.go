package day15

// SpokenNumbers finds the last spoken number given
// the turns to run
func SpokenNumbers(firstNumbers []int, turns int) int {
	numbers := make([]int, turns)
	spoken := make(map[int][]int)
	copy(numbers, firstNumbers)
	for i := range numbers {
		if i < len(firstNumbers) {
			spoken[firstNumbers[i]] = append(spoken[firstNumbers[i]], i+1)
			continue
		}
		n := numbers[i-1]
		prev, ok := spoken[n]
		if !ok {
			numbers[i] = 0
			spoken[0] = append(spoken[0], i+1)
		} else {
			var result int
			if len(prev) == 1 {
				result = i - prev[0]
			} else {
				result = prev[len(prev)-1] - prev[len(prev)-2]
			}
			numbers[i] = result
			spoken[result] = append(spoken[result], i+1)
		}
	}
	return numbers[turns-1]
}

// EfficientSpokenNumbers finds the spoken number given the turns to run
// in a more efficient way (half time compared with previous one)
func EfficientSpokenNumbers(firstNumbers []int, turns int) int {
	a := make([]int, turns) // all zero by default
	copy(a, firstNumbers)
	var (
		seen = make(map[int]int)
		m    int
		ok   bool
	)
	for n := 0; n < turns-1; n++ {
		if m, ok = seen[a[n]]; ok {
			a[n+1] = n - m
		}
		seen[a[n]] = n
	}
	return a[turns-1]
}
