package day3

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
			name: "test case",
			path: "./test.txt",
			want: 7,
		},
		{
			name: "real part 1",
			path: "./input.txt",
			want: 164,
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
			name: "test case",
			path: "./test.txt",
			want: 336,
		},
		{
			name: "real part 1",
			path: "./input.txt",
			want: 5007658656,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := SecondPart(lines); got != tt.want {
				t.Errorf("SecondPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
