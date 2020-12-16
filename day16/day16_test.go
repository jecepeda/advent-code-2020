package day16

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
			want: 71,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 26869,
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
		name    string
		path    string
		want    int
		keyword string
	}{
		{
			name:    "test",
			path:    "./test.txt",
			want:    14,
			keyword: "seat",
		},
		{
			name:    "input",
			path:    "./input.txt",
			keyword: "departure",
			want:    855275529001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			got, err := SecondPart(lines, tt.keyword)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
