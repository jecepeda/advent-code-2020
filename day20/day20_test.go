package day20

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
			want: 20899048083289,
		},
		{
			name: "input",
			path: "./input.txt",
			want: 7492183537913,
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
