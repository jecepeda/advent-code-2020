package day22

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
			want: 306,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 34324,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got, err := FirstPart(lines)
			require.NoError(t, err)
			require.Equal(t, tr.want, got)
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
			want: 291,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 33259,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got, err := SecondPart(lines)
			require.NoError(t, err)
			require.Equal(t, tr.want, got)
		})
	}
}

func TestRecursiveCombat(t *testing.T) {
	got, _ := playRecursiveCombat([]int{43, 19}, []int{2, 29, 14})
	require.Equal(t, 1, got)
}
