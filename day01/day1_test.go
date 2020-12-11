package day01

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/require"
)

func TestSum2020Part1(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     int
	}{
		{
			name:     "test case",
			filePath: "./test.txt",
			want:     514579,
		},
		{
			name:     "first case",
			filePath: "./input.txt",
			want:     436404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := util.FileToIntList(tt.filePath)
			require.NoError(t, err)
			if got := Sum2020Part1(numbers); got != tt.want {
				t.Errorf("Sum2020Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum2020Part2(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     int
	}{
		{
			name:     "test case",
			filePath: "./test.txt",
			want:     241861950,
		},
		{
			name:     "first case",
			filePath: "./input.txt",
			want:     274879808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := util.FileToIntList(tt.filePath)
			require.NoError(t, err)
			if got := Sum2020Part2(numbers); got != tt.want {
				t.Errorf("Sum2020Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
