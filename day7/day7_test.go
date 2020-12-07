package day7

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBagMatcher(t *testing.T) {
	b, err := NewBagMatcher()
	require.NoError(t, err)
	line := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	color := b.GetBagColor(line)
	assert.Equal(t, "light red", color)
	assert.True(t, b.HasBags(line))
	matches, err := b.GetBagsContained(line)
	require.NoError(t, err)
	require.Len(t, matches, 2)
}

func TestFirstPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test",
			path: "./test.txt",
			want: 4,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 144,
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
			want: 32,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 5956,
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
