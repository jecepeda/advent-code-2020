package day3

// FirstPart checks the number of trees that are found in
// our way
// the path we follow is
// 3 right 1 down
func FirstPart(lines []string) int {
	return TraverseForest(lines, 3, 1)

}

// SecondPart checks the number of trees that are found in our way
// we follow different parts and multiply the number of trees found
// we follow this order
// 1-1, 3-1, 5-1, 7-1, 1-2
func SecondPart(lines []string) int {
	oneOne := TraverseForest(lines, 1, 1)
	threeOne := TraverseForest(lines, 3, 1)
	fiveOne := TraverseForest(lines, 5, 1)
	sevenOne := TraverseForest(lines, 7, 1)
	oneTwo := TraverseForest(lines, 1, 2)
	return oneOne * threeOne * fiveOne * sevenOne * oneTwo
}

// TraverseForest traverse the forest in the given pattern, returning
// the number of trees found on the way
func TraverseForest(lines []string, right int, down int) int {
	var (
		nOfTrees int
		i, j     int
	)
	for i < len(lines) {
		if lines[i][j] == '#' {
			nOfTrees++
		}
		j += right
		j %= len(lines[i])
		i += down
	}
	return nOfTrees
}
