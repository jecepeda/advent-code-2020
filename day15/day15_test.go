package day15

import "testing"

func TestSpokenNumbersFirstPart(t *testing.T) {
	type args struct {
		firstNumbers []int
		turns        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				firstNumbers: []int{0, 3, 6},
				turns:        2020,
			},
			want: 436,
		},
		{
			name: "example 2",
			args: args{
				firstNumbers: []int{1, 3, 2},
				turns:        2020,
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				firstNumbers: []int{2, 1, 3},
				turns:        2020,
			},
			want: 10,
		},
		{
			name: "example 4",
			args: args{
				firstNumbers: []int{1, 2, 3},
				turns:        2020,
			},
			want: 27,
		},
		{
			name: "example 5",
			args: args{
				firstNumbers: []int{2, 3, 1},
				turns:        2020,
			},
			want: 78,
		},
		{
			name: "example 6",
			args: args{
				firstNumbers: []int{3, 2, 1},
				turns:        2020,
			},
			want: 438,
		},
		{
			name: "example 7",
			args: args{
				firstNumbers: []int{3, 1, 2},
				turns:        2020,
			},
			want: 1836,
		},
		{
			name: "input case",
			args: args{
				firstNumbers: []int{0, 14, 1, 3, 7, 9},
				turns:        2020,
			},
			want: 763,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpokenNumbers(tt.args.firstNumbers, tt.args.turns); got != tt.want {
				t.Errorf("FirstPart: SpokenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpokenNumbersSecondPart(t *testing.T) {
	type args struct {
		firstNumbers []int
		turns        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				firstNumbers: []int{0, 3, 6},
				turns:        30000000,
			},
			want: 175594,
		},
		{
			name: "example 2",
			args: args{
				firstNumbers: []int{1, 3, 2},
				turns:        30000000,
			},
			want: 2578,
		},
		{
			name: "example 3",
			args: args{
				firstNumbers: []int{2, 1, 3},
				turns:        30000000,
			},
			want: 3544142,
		},
		{
			name: "example 4",
			args: args{
				firstNumbers: []int{1, 2, 3},
				turns:        30000000,
			},
			want: 261214,
		},
		{
			name: "example 5",
			args: args{
				firstNumbers: []int{2, 3, 1},
				turns:        30000000,
			},
			want: 6895259,
		},
		{
			name: "example 6",
			args: args{
				firstNumbers: []int{3, 2, 1},
				turns:        30000000,
			},
			want: 18,
		},
		{
			name: "example 7",
			args: args{
				firstNumbers: []int{3, 1, 2},
				turns:        30000000,
			},
			want: 362,
		},
		{
			name: "input case",
			args: args{
				firstNumbers: []int{0, 14, 1, 3, 7, 9},
				turns:        30000000,
			},
			want: 1876406,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EfficientSpokenNumbers(tt.args.firstNumbers, tt.args.turns); got != tt.want {
				t.Errorf("SecondPart: SpokenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
