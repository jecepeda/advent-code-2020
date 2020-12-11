package day11

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/require"
)

func TestFirstPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test",
			path: "./test.txt",
			want: 37,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 2494,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			got := FirstPart(lines)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSecondPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test",
			path: "./test.txt",
			want: 26,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 2306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			got := SecondPart(lines)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_numberOfOccupiedSeats(t *testing.T) {
	seatMap := [][]rune{
		{'#', '#', '#', '#'},
		{'#', '#', '#', '#'},
		{'#', '#', '#', '#'},
		{'#', '#', '#', '#'},
	}
	type args struct {
		lines [][]rune
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 0 case",
			args: args{
				lines: seatMap,
				i:     0,
				j:     0,
			},
			want: 3,
		},
		{
			name: "0 1 case",
			args: args{
				lines: seatMap,
				i:     0,
				j:     1,
			},
			want: 5,
		},
		{
			name: "1 0 case",
			args: args{
				lines: seatMap,
				i:     1,
				j:     0,
			},
			want: 5,
		},
		{
			name: "2 2 case",
			args: args{
				lines: seatMap,
				i:     2,
				j:     2,
			},
			want: 8,
		},
		{
			name: "3 3 case",
			args: args{
				lines: seatMap,
				i:     3,
				j:     3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfAdjacentOccupiedSeats(tt.args.lines, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("numberOfOccupiedSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfVisibleOccupiedSeats(t *testing.T) {
	eightMap := [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', 'L', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}
	oneFreeSeat := [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'L', '.', 'L', '.', '#', '.', '#', '.', '#', '.', '#', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	noOccupied := [][]rune{
		{'.', '#', '#', '.', '#', '#', '.'},
		{'#', '.', '#', '.', '#', '.', '#'},
		{'#', '#', '.', '.', '.', '#', '#'},
		{'.', '.', '.', 'L', '.', '.', '.'},
		{'#', '#', '.', '.', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '#'},
		{'.', '#', '#', '.', '#', '#', '.'},
	}
	type args struct {
		lines [][]rune
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "eight seats",
			args: args{
				lines: eightMap,
				i:     4,
				j:     3,
			},
			want: 8,
		},
		{
			name: "one free seat",
			args: args{
				lines: oneFreeSeat,
				i:     1,
				j:     1,
			},
			want: 0,
		},
		{
			name: "no occupied seat",
			args: args{
				lines: noOccupied,
				i:     3,
				j:     3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfVisibleOccupiedSeats(tt.args.lines, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("numberOfVisibleOccupiedSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
