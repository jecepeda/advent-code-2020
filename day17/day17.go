package day17

// values that Conway's cube can get
const (
	ACTIVE   = '#'
	INACTIVE = '.'
)

// Program represents a Conway's cube program
type Program struct {
	Map3D  [][][]rune
	Map4D  [][][][]rune
	Rounds int
}

// New3DProgram creates a new Program for a 3d map
func New3DProgram(m [][][]rune, rounds int) *Program {
	return &Program{
		Map3D:  m,
		Rounds: rounds,
	}
}

// New4DProgram creates a new program for a 4d map
func New4DProgram(m [][][][]rune, rounds int) *Program {
	return &Program{
		Map4D:  m,
		Rounds: rounds,
	}
}

// Adjacent3DActives tries to find the actives adjacent to this position
func (p *Program) Adjacent3DActives(x, y, z int) int {
	actives := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if (i >= 0 && i < len(p.Map3D)) && (j >= 0 && j < len(p.Map3D[i])) && (k >= 0 && k < len(p.Map3D[i][j])) {
					if (i != x || j != y || k != z) && p.Map3D[i][j][k] == ACTIVE {
						actives++
					}
				}
			}
		}
	}
	return actives
}

// Adjacent4DActives tries to find the actives adjacent to this position
func (p *Program) Adjacent4DActives(x, y, z, w int) int {
	actives := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				for l := w - 1; l <= w+1; l++ {
					if (i >= 0 && i < len(p.Map4D)) && (j >= 0 && j < len(p.Map4D[i])) && (k >= 0 && k < len(p.Map4D[i][j])) && (l >= 0 && l < len(p.Map4D[i][j][k])) {
						if (i != x || j != y || k != z || l != w) && p.Map4D[i][j][k][l] == ACTIVE {
							actives++
						}
					}
				}
			}
		}
	}
	return actives
}

// Run3D runs the program
func (p *Program) Run3D() {
	for i := 0; i < p.Rounds; i++ {
		newMap := copy3DMap(p.Map3D)
		for i := range p.Map3D {
			for j := range p.Map3D[i] {
				for k := range p.Map3D[i][j] {
					v := p.Adjacent3DActives(i, j, k)
					if p.Map3D[i][j][k] == ACTIVE {
						if v == 2 || v == 3 {
							newMap[i][j][k] = ACTIVE
						} else {
							newMap[i][j][k] = INACTIVE
						}
					} else if p.Map3D[i][j][k] == INACTIVE {
						if v == 3 {
							newMap[i][j][k] = ACTIVE
						}
					}
				}
			}
		}
		p.Map3D = newMap
	}
}

// Run4D runs the program with a 4 dimensional matrix
func (p *Program) Run4D() {
	for i := 0; i < p.Rounds; i++ {
		newMap := copy4DMap(p.Map4D)
		for i := range p.Map4D {
			for j := range p.Map4D[i] {
				for k := range p.Map4D[i][j] {
					for l := range p.Map4D[i][j][k] {
						v := p.Adjacent4DActives(i, j, k, l)
						if p.Map4D[i][j][k][l] == ACTIVE {
							if v == 2 || v == 3 {
								newMap[i][j][k][l] = ACTIVE
							} else {
								newMap[i][j][k][l] = INACTIVE
							}
						} else if p.Map4D[i][j][k][l] == INACTIVE {
							if v == 3 {
								newMap[i][j][k][l] = ACTIVE
							}
						}
					}
				}
			}
		}
		p.Map4D = newMap
	}
}

// Count3DActives counts the number of active positions in
// the whole conway's map
func (p *Program) Count3DActives() int {
	actives := 0
	for i := range p.Map3D {
		for j := range p.Map3D[i] {
			for k := range p.Map3D[i][j] {
				if p.Map3D[i][j][k] == ACTIVE {
					actives++
				}
			}
		}
	}
	return actives
}

// Count4DActives counts the number of active positions in
// the whole conway's map
func (p *Program) Count4DActives() int {
	actives := 0
	for i := range p.Map4D {
		for j := range p.Map4D[i] {
			for k := range p.Map4D[i][j] {
				for l := range p.Map4D[i][j][k] {
					if p.Map4D[i][j][k][l] == ACTIVE {
						actives++
					}
				}
			}
		}
	}
	return actives
}

func copy3DMap(m [][][]rune) [][][]rune {
	newRune := make([][][]rune, len(m))
	for i := range newRune {
		newRune[i] = make([][]rune, len(m[i]))
		for j := range newRune[i] {
			newRune[i][j] = make([]rune, len(m[i][j]))
			copy(newRune[i][j], m[i][j])
		}
	}
	return newRune
}

func copy4DMap(m [][][][]rune) [][][][]rune {
	newRune := make([][][][]rune, len(m))
	for i := range newRune {
		newRune[i] = make([][][]rune, len(m[i]))
		for j := range newRune[i] {
			newRune[i][j] = make([][]rune, len(m[i][j]))
			for k := range newRune[i][j] {
				newRune[i][j][k] = make([]rune, len(m[i][j][k]))
				copy(newRune[i][j][k], m[i][j][k])
			}
		}
	}
	return newRune
}

func createMultiplied3DMap(size int, origMap [][]rune) [][][]rune {
	newSize := size * 5
	m := make([][][]rune, newSize)
	for i := range m {
		m[i] = make([][]rune, newSize)
		for j := range m[i] {
			m[i][j] = make([]rune, newSize)
		}
	}
	middle := newSize / 2
	for i := range origMap {
		for j := range origMap[i] {
			m[middle][middle+i][middle+j] = origMap[i][j]
		}
	}
	// paint empties to inactive
	for i := range m {
		for j := range m[i] {
			for k := range m[i][j] {
				if m[i][j][k] != ACTIVE {
					m[i][j][k] = INACTIVE
				}
			}
		}
	}

	return m
}

func createMultiplied4DMap(size int, origMap [][]rune) [][][][]rune {
	newSize := size * 5
	m := make([][][][]rune, newSize)
	for i := range m {
		m[i] = make([][][]rune, newSize)
		for j := range m[i] {
			m[i][j] = make([][]rune, newSize)
			for k := range m[i][j] {
				m[i][j][k] = make([]rune, newSize)
			}
		}
	}
	middle := newSize / 2
	for i := range origMap {
		for j := range origMap[i] {
			m[middle][middle][middle+i][middle+j] = origMap[i][j]
		}
	}
	// paint empties to inactive
	for i := range m {
		for j := range m[i] {
			for k := range m[i][j] {
				for l := range m[i][j][k] {
					if m[i][j][k][l] != ACTIVE {
						m[i][j][k][l] = INACTIVE
					}
				}
			}
		}
	}

	return m
}

// FirstPart creates a conway cube, runs a number of
// rounds and returns the number of active cells
func FirstPart(lines []string, rounds int) int {
	firstMap := make([][]rune, len(lines))
	for i := range firstMap {
		firstMap[i] = []rune(lines[i])
	}
	mMap := createMultiplied3DMap(len(firstMap), firstMap)
	p := New3DProgram(mMap, rounds)
	p.Run3D()
	return p.Count3DActives()
}

// SecondPart creates a conway 4d cube, runs a number of
// rounds and returns the number of active cells
func SecondPart(lines []string, rounds int) int {
	firstMap := make([][]rune, len(lines))
	for i := range firstMap {
		firstMap[i] = []rune(lines[i])
	}
	mMap := createMultiplied4DMap(len(firstMap), firstMap)
	p := New4DProgram(mMap, rounds)
	p.Run4D()
	return p.Count4DActives()
}
