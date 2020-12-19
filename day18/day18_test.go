package day18

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
			name: "input",
			path: "./input.txt",
			want: 800602729153,
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
		name   string
		path   string
		rounds int
		want   int
	}{
		{
			name: "input",
			path: "./input.txt",
			want: 92173009047076,
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

func TestFirstPartNoFile(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "example case 1",
			args:    args{lines: []string{"2 * 3 + (4 * 5)"}},
			want:    26,
			wantErr: false,
		},
		{
			name:    "example case 2",
			args:    args{lines: []string{"5 + (8 * 3 + 9 + 3 * 4 * 3"}},
			want:    437,
			wantErr: false,
		},
		{
			name:    "example case 3",
			args:    args{lines: []string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}},
			want:    12240,
			wantErr: false,
		},
		{
			name:    "example case 4",
			args:    args{lines: []string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}},
			want:    13632,
			wantErr: false,
		},
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstPart(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSecondPartNoFile(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "example case 1",
			args:    args{lines: []string{"2 * 3 + (4 * 5)"}},
			want:    46,
			wantErr: false,
		},
		{
			name:    "example case 2",
			args:    args{lines: []string{"5 + (8 * 3 + 9 + 3 * 4 * 3"}},
			want:    1445,
			wantErr: false,
		},
		{
			name:    "example case 3",
			args:    args{lines: []string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}},
			want:    669060,
			wantErr: false,
		},
		{
			name:    "example case 4",
			args:    args{lines: []string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}},
			want:    23340,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecondPart(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compute(t *testing.T) {
	type args struct {
		q *Queue
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple",
			args: args{
				q: &Queue{
					Data: []Token{
						{Kind: Number, Number: 2},
						{Kind: Mul},
						{Kind: Number, Number: 2},
					},
				},
			},
			want: 4,
		},
		{
			name: "more than 1 operation",
			args: args{
				q: &Queue{
					Data: []Token{
						{Kind: Number, Number: 2},
						{Kind: Sum},
						{Kind: Number, Number: 2},
						{Kind: Mul},
						{Kind: Number, Number: 2},
					},
				},
			},
			want: 8,
		},
		{
			name: "parenthesis",
			args: args{
				q: &Queue{
					Data: []Token{
						{Kind: LeftParen},
						{Kind: Number, Number: 2},
						{Kind: Sum},
						{Kind: Number, Number: 2},
						{Kind: RightParen},
					},
				},
			},
			want: 4,
		},
		{
			name: "more parenthesis",
			args: args{
				q: &Queue{
					Data: []Token{
						{Kind: LeftParen},
						{Kind: Number, Number: 2},
						{Kind: Sum},
						{Kind: Number, Number: 2},
						{Kind: RightParen},
						{Kind: Mul},
						{Kind: Number, Number: 2},
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.q); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRPNQueue(t *testing.T) {
	tests := []struct {
		name string
		args []Token
		want Queue
	}{
		{
			name: "simple",
			args: []Token{
				{Kind: Number, Number: 2},
				{Kind: Mul},
				{Kind: Number, Number: 3},
			},
			want: Queue{
				Data: []Token{
					{Kind: Number, Number: 2},
					{Kind: Number, Number: 3},
					{Kind: Mul},
				},
			},
		},
		{
			name: "more complex",
			args: []Token{
				{Kind: Number, Number: 2},
				{Kind: Mul},
				{Kind: Number, Number: 3},
				{Kind: Sum},
				{Kind: Number, Number: 4},
			},
			want: Queue{
				Data: []Token{
					{Kind: Number, Number: 2},
					{Kind: Number, Number: 3},
					{Kind: Number, Number: 4},
					{Kind: Sum},
					{Kind: Mul},
				},
			},
		},
		{
			name: "with parenthesis",
			args: []Token{
				{Kind: LeftParen},
				{Kind: Number, Number: 2},
				{Kind: Mul},
				{Kind: Number, Number: 3},
				{Kind: RightParen},
				{Kind: Sum},
				{Kind: Number, Number: 4},
			},
			want: Queue{
				Data: []Token{
					{Kind: Number, Number: 2},
					{Kind: Number, Number: 3},
					{Kind: Mul},
					{Kind: Number, Number: 4},
					{Kind: Sum},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRPNQueue(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
