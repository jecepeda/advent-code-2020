package day12

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
			want: 25,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 1482,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			got, err := FirstPart(lines)
			require.NoError(t, err)
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
			want: 286,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 48739,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			got, err := SecondPart(lines)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
