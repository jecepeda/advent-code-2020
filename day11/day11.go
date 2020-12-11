package day11

import (
	"reflect"
)

// values input can take
const (
	Floor    = '.'
	Free     = 'L'
	Occupied = '#'
)

// FirstPart applies the rules of seat changing
// and returns the number of occupied seats once they stabilize
func FirstPart(lines []string) int {
	prev := make([][]rune, len(lines))
	for idx, l := range lines {
		prev[idx] = []rune(l)
	}
	for {
		seatMap := changeAdjacentSeats(prev)
		if reflect.DeepEqual(seatMap, prev) {
			return countOccupiedSeats(seatMap)
		}
		prev = seatMap
	}
}

// SecondPart applies the rules of seat changing
// and returns the number of occupied seats once they stabilize
func SecondPart(lines []string) int {
	prev := make([][]rune, len(lines))
	for idx, l := range lines {
		prev[idx] = []rune(l)
	}
	for {
		seatMap := changeVisibleSeats(prev)
		if reflect.DeepEqual(seatMap, prev) {
			return countOccupiedSeats(seatMap)
		}
		prev = seatMap
	}
}

func copyMap(seatMap [][]rune) [][]rune {
	newSeatMap := make([][]rune, len(seatMap))
	for i := range seatMap {
		newSeatMap[i] = make([]rune, len(seatMap[i]))
		copy(newSeatMap[i], seatMap[i])
	}
	return newSeatMap
}

func countOccupiedSeats(seatMap [][]rune) int {
	var occupied int
	for i := range seatMap {
		for j := range seatMap[i] {
			if seatMap[i][j] == Occupied {
				occupied++
			}
		}
	}
	return occupied
}

func changeAdjacentSeats(seatMap [][]rune) [][]rune {
	newMap := copyMap(seatMap)
	for i := range seatMap {
		for j := range seatMap[i] {
			switch seatMap[i][j] {
			case Occupied:
				if numberOfAdjacentOccupiedSeats(seatMap, i, j) > 3 {
					newMap[i][j] = Free
				}
			case Free:
				if numberOfAdjacentOccupiedSeats(seatMap, i, j) == 0 {
					newMap[i][j] = Occupied
				}
			}
		}
	}
	return newMap
}

func changeVisibleSeats(seatMap [][]rune) [][]rune {
	newMap := copyMap(seatMap)
	for i := range seatMap {
		for j := range seatMap[i] {
			switch seatMap[i][j] {
			case Occupied:
				if numberOfVisibleOccupiedSeats(seatMap, i, j) > 4 {
					newMap[i][j] = Free
				}
			case Free:
				if numberOfVisibleOccupiedSeats(seatMap, i, j) == 0 {
					newMap[i][j] = Occupied
				}
			}
		}
	}
	return newMap
}

func numberOfVisibleOccupiedSeats(seatMap [][]rune, i, j int) int {
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, 0}, {1, 1}, {1, -1}}
	var occupied int
	for _, direction := range directions {
		x, y := i, j
		for x+direction[0] >= 0 && x+direction[0] < len(seatMap) && y+direction[1] >= 0 && y+direction[1] < len(seatMap[0]) {
			x += direction[0]
			y += direction[1]
			if seatMap[x][y] == Occupied {
				occupied++
				break
			} else if seatMap[x][y] == Free {
				break
			}
		}
	}
	return occupied
}

func numberOfAdjacentOccupiedSeats(seatMap [][]rune, i, j int) int {
	lookups := make([][2]int, 0)
	if i < len(seatMap)-1 {
		lookups = append(lookups, [2]int{i + 1, j})
		if j < len(seatMap[i])-1 {
			lookups = append(lookups, [2]int{i + 1, j + 1})
		}
		if j > 0 {
			lookups = append(lookups, [2]int{i + 1, j - 1})
		}
	}
	if i > 0 {
		lookups = append(lookups, [2]int{i - 1, j})
		if j > 0 {
			lookups = append(lookups, [2]int{i - 1, j - 1})
		}
		if j < len(seatMap[i])-1 {
			lookups = append(lookups, [2]int{i - 1, j + 1})
		}
	}
	if j < len(seatMap[i])-1 {
		lookups = append(lookups, [2]int{i, j + 1})
	}
	if j > 0 {
		lookups = append(lookups, [2]int{i, j - 1})
	}
	occupied := 0
	for _, lookup := range lookups {
		if seatMap[lookup[0]][lookup[1]] == Occupied {
			occupied++
		}
	}
	return occupied
}
