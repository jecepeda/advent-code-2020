package day05

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoardingPass_Decode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected *BoardingPass
	}{
		{
			name: "example case",
			code: "FBFBBFFRLR",
			expected: &BoardingPass{
				Code: "FBFBBFFRLR",
				Row:  44,
				Seat: 5,
				ID:   357,
			},
		},
		{
			name: "example case 2",
			code: "BFFFBBFRRR",
			expected: &BoardingPass{
				Code: "BFFFBBFRRR",
				Row:  70,
				Seat: 7,
				ID:   567,
			},
		},
		{
			name: "example case 3",
			code: "BBFFBBFRLL",
			expected: &BoardingPass{
				Code: "BBFFBBFRLL",
				Row:  102,
				Seat: 4,
				ID:   820,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BoardingPass{
				Code: tt.code,
			}
			b.Decode()
			assert.Equal(t, tt.expected, b)
		})
	}
}

func TestFirstPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "input",
			path: "./input.txt",
			want: 885,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := FirstPart(lines); got != tt.want {
				t.Errorf("FirstPart() = %v, want %v", got, tt.want)
			}
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
			name: "input",
			path: "./input.txt",
			want: 623,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := SecondPart(lines); got != tt.want {
				t.Errorf("FirstPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
