package day10

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
			name: "test 1",
			path: "./test_1.txt",
			want: 35,
		},
		{
			name: "test 2",
			path: "./test_2.txt",
			want: 220,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 2080,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := util.FileToIntList(tt.path)
			require.NoError(t, err)
			got := FirstPart(numbers)
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
			name: "test 1",
			path: "./test_1.txt",
			want: 8,
		},
		{
			name: "test 2",
			path: "./test_2.txt",
			want: 19208,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 6908379398144,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := util.FileToIntList(tt.path)
			require.NoError(t, err)
			got := SecondPart(numbers)
			require.Equal(t, tt.want, got)
		})
	}
}
