package day6

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
			want: 11,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 6161,
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
			name: "test",
			path: "./test.txt",
			want: 6,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 2971,
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
