package day21

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
			want: 5,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 2724,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got := FirstPart(lines)
			require.Equal(t, tr.want, got)
		})
	}
}

func TestSecondPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "test",
			path: "./test.txt",
			want: "mxmxvkd,sqjhc,fvjkl",
		},
		{
			name: "input",
			path: "./input.txt",
			want: "xlxknk,cskbmx,cjdmk,bmhn,jrmr,tzxcmr,fmgxh,fxzh",
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			lines, err := util.ReadFile(tr.path)
			require.NoError(t, err)
			got := SecondPart(lines)
			require.Equal(t, tr.want, got)
		})
	}
}
