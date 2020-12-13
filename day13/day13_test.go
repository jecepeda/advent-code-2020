package day13

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
			want: 295,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 2406,
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
			want: 1068781,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 225850756401039,
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

func TestSecondPartManual(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "test 1",
			want: 3417,
			lines: []string{
				"",
				"17,x,13,19",
			},
		},
		{
			name: "test 2",
			want: 754018,
			lines: []string{
				"",
				"67,7,59,61",
			},
		},
		{
			name: "test 3",
			want: 779210,
			lines: []string{
				"",
				"67,x,7,59,61",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecondPart(tt.lines)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
