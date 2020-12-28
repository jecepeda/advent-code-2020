package day19

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
			want: 2,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 178,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got, err := ValidateSentences(lines)
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
			path: "./test_replacements.txt",
			want: 12,
		},
		{
			name: "input",
			path: "./input_replacements.txt",
			want: 346,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got, err := ValidateSentences(lines)
			require.NoError(t, err)
			require.Equal(t, tr.want, got)
		})
	}
}
