package day2

import (
	"testing"

	"github.com/jecepeda/advent-code-2020/util"
	"github.com/stretchr/testify/require"
)

func TestCheckPasswordsPart1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test case",
			path: "./test.txt",
			want: 2,
		},
		{
			name: "real part 1",
			path: "./input.txt",
			want: 434,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := CheckPasswordsPart1(lines); got != tt.want {
				t.Errorf("CheckPasswordsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswordsPart2(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{
			name: "test case",
			path: "./test.txt",
			want: 1,
		},
		{
			name: "real part 1",
			path: "./input.txt",
			want: 509,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := util.ReadFile(tt.path)
			require.NoError(t, err)
			if got := CheckPasswordsPart2(lines); got != tt.want {
				t.Errorf("CheckPasswordsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
