package day09

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/require"
)

func TestFindNotMatchingNumber(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		preamble int
		want     int
	}{
		{
			name:     "test",
			path:     "./test.txt",
			want:     127,
			preamble: 5,
		},
		{
			name:     "input",
			path:     "./input.txt",
			want:     31161678,
			preamble: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := util.FileToIntList(tt.path)
			require.NoError(t, err)
			got := FindNotMatchingNumber(numbers, tt.preamble)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSecondPart(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		preamble int
		want     int
	}{
		{
			name:     "test",
			path:     "./test.txt",
			want:     62,
			preamble: 5,
		},
		{
			name:     "input",
			path:     "./input.txt",
			want:     5453868,
			preamble: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := util.FileToIntList(tt.path)
			require.NoError(t, err)
			got := SecondPart(numbers, tt.preamble)
			require.Equal(t, tt.want, got)
		})
	}
}
