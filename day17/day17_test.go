package day17

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/require"
)

func TestFirstPart(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		rounds int
		want   int
	}{
		{
			name:   "test one round",
			path:   "./test.txt",
			rounds: 6,
			want:   109,
		},
		{
			name:   "input",
			path:   "./input.txt",
			rounds: 6,
			want:   213,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got := FirstPart(lines, tr.rounds)
			require.Equal(t, tr.want, got)
		})
	}
}

func TestSecondPart(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		rounds int
		want   int
	}{
		{
			name:   "test one round",
			path:   "./test.txt",
			rounds: 6,
			want:   800,
		},
		{
			name:   "input",
			path:   "./input.txt",
			rounds: 6,
			want:   1624,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got := SecondPart(lines, tr.rounds)
			require.Equal(t, tr.want, got)
		})
	}
}
