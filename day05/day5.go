package day05

// BoardingPass is the binary-space partitioning
// pass used to know where you should seat
type BoardingPass struct {
	Code string
	Row  int
	Seat int
	ID   int
}

// NewBoardingPass generates a new boarding pass
func NewBoardingPass(str string) *BoardingPass {
	return &BoardingPass{
		Code: str,
	}
}

// Decode decodes the boarding pass, getting
// the row, seat, and the id of the boarding pass
func (b *BoardingPass) Decode() {
	b.Row = binarySearch(b.Code[:7], 0, 0, 127)
	b.Seat = binarySearch(b.Code[7:], 0, 0, 7)
	b.ID = (b.Row * 8) + b.Seat
}

func binarySearch(str string, pos, lower, upper int) int {
	if pos == (len(str)) {
		return upper
	}
	middle := (lower + upper) / 2
	if str[pos] == 'F' || str[pos] == 'L' {
		return binarySearch(str, pos+1, lower, middle)
	}
	return binarySearch(str, pos+1, middle+1, upper)
}

// FirstPart looks into the boarding passes and returns the highest one
func FirstPart(lines []string) int {
	highest := 0
	for _, l := range lines {
		bp := NewBoardingPass(l)
		bp.Decode()
		if bp.ID > highest {
			highest = bp.ID
		}
	}
	return highest
}

// SecondPart looks into the boarding passes
// and checks which is your seat
func SecondPart(lines []string) int {
	boardingpasses := make(map[int]*BoardingPass)
	lowest, highest := 999999999, 0
	var ok, okLower, okHigher bool
	for _, l := range lines {
		bp := NewBoardingPass(l)
		bp.Decode()
		if bp.ID > highest {
			highest = bp.ID
		}
		if bp.ID < lowest {
			lowest = bp.ID
		}
		boardingpasses[bp.ID] = bp
	}
	for i := lowest; i < highest; i++ {
		_, ok = boardingpasses[i]
		if ok {
			continue
		}
		_, okLower = boardingpasses[i-1]
		_, okHigher = boardingpasses[i+1]
		if okLower && okHigher {
			return i
		}
	}
	return 0
}
