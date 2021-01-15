package day23

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle_PickFrom(t *testing.T) {
	type args struct {
		pos int
	}
	tests := []struct {
		name   string
		fields []int
		pos    int
		res    []int
		want   *Circle
	}{
		{
			name:   "base",
			fields: []int{1, 2, 3, 4, 5, 6},
			pos:    0,
			res:    []int{2, 3, 4},
			want: &Circle{
				variables: []int{1, 5, 6},
			},
		},
		{
			name:   "complex case",
			fields: []int{1, 2, 3, 4, 5, 6},
			pos:    4,
			res:    []int{6, 1, 2},
			want: &Circle{
				variables: []int{3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				variables: tt.fields,
			}
			if got := c.PickFrom(tt.pos); !reflect.DeepEqual(got, tt.res) {
				t.Errorf("Circle.PickFrom() = %v, want %v", got, tt.res)
			}
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Circle() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestCircle_Insert(t *testing.T) {
	type fields struct {
		variables []int
	}
	type args struct {
		value    int
		valuePos int
		values   []int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *Circle
	}{
		{
			name:   "1",
			fields: fields{variables: []int{3, 2, 5, 4, 6, 7}},
			args: args{
				value:    3,
				valuePos: 0,
				values:   []int{8, 9, 1},
			},
			expected: &Circle{
				variables: []int{3, 2, 8, 9, 1, 5, 4, 6, 7},
			},
		},
		{
			name: "2",
			fields: fields{
				variables: []int{3, 2, 5, 4, 6, 7},
			},
			args: args{
				value:    2,
				valuePos: 1,
				values:   []int{8, 9, 1},
			},
			expected: &Circle{
				variables: []int{3, 2, 5, 4, 6, 7, 8, 9, 1},
			},
		},
		{
			name: "3",
			fields: fields{
				variables: []int{3, 2, 5, 8, 9, 1},
			},
			args: args{
				value:    5,
				valuePos: 2,
				values:   []int{4, 6, 7},
			},
			expected: &Circle{
				variables: []int{7, 2, 5, 8, 9, 1, 3, 4, 6},
			},
		},
		{
			name: "4",
			fields: fields{
				variables: []int{7, 2, 5, 8, 4, 6},
			},
			args: args{
				value:    8,
				valuePos: 3,
				values:   []int{9, 1, 3},
			},
			expected: &Circle{
				variables: []int{3, 2, 5, 8, 4, 6, 7, 9, 1},
			},
		},
		{
			name: "5",
			fields: fields{
				variables: []int{3, 2, 5, 8, 4, 1},
			},
			args: args{
				value:    4,
				valuePos: 4,
				values:   []int{6, 7, 9},
			},
			expected: &Circle{
				variables: []int{9, 2, 5, 8, 4, 1, 3, 6, 7},
			},
		},
		{
			name: "6",
			fields: fields{
				variables: []int{9, 2, 5, 8, 4, 1},
			},
			args: args{
				value:    1,
				valuePos: 5,
				values:   []int{3, 6, 7},
			},
			expected: &Circle{
				variables: []int{7, 2, 5, 8, 4, 1, 9, 3, 6},
			},
		},
		{
			name: "7",
			fields: fields{
				variables: []int{2, 5, 8, 4, 1, 9},
			},
			args: args{
				value:    9,
				valuePos: 6,
				values:   []int{3, 6, 7},
			},
			expected: &Circle{
				variables: []int{8, 3, 6, 7, 4, 1, 9, 2, 5},
			},
		},
		{
			name: "8",
			fields: fields{
				variables: []int{6, 7, 4, 1, 9, 2},
			},
			args: args{
				value:    2,
				valuePos: 7,
				values:   []int{5, 8, 3},
			},
			expected: &Circle{
				variables: []int{7, 4, 1, 5, 8, 3, 9, 2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				variables: tt.fields.variables,
			}
			c.Insert(tt.args.value, tt.args.valuePos, tt.args.values)
			assert.Equal(t, tt.expected, c)
		})
	}
}

func TestFirstPart(t *testing.T) {
	type args struct {
		values []int
		rounds int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example case",
			args: args{
				values: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				rounds: 10,
			},
			want: "92658374",
		},
		{
			name: "example case 2",
			args: args{
				values: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				rounds: 100,
			},
			want: "67384529",
		},
		{
			name: "real case",
			args: args{
				values: []int{3, 9, 4, 6, 1, 8, 5, 2, 7},
				rounds: 100,
			},
			want: "78569234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstPart(tt.args.values, tt.args.rounds); got != tt.want {
				t.Errorf("FirstPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecondPart(t *testing.T) {
	type args struct {
		values []int
		size   int
		rounds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example case",
			args: args{
				values: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				size:   1_000_000,
				rounds: 10_000_000,
			},
			want: 149245887792,
		},
		{
			name: "real case",
			args: args{
				values: []int{3, 9, 4, 6, 1, 8, 5, 2, 7},
				size:   1_000_000,
				rounds: 10_000_000,
			},
			want: 565615814504,
		},
	}
	for _, tt := range tests {
		tr := tt
		t.Run(tr.name, func(t *testing.T) {
			t.Parallel()
			if got := SecondPart(tr.args.values, tr.args.size, tr.args.rounds); got != tr.want {
				t.Errorf("SecondPart() = %v, want %v", got, tr.want)
			}
		})
	}
}
