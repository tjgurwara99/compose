package compose

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		input  []int
		filter func(int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple",
			args: args{
				input:  []int{1, 2},
				filter: strconv.Itoa,
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.input, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		input     []int
		predicate func(item int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple case",
			args: args{
				input:     []int{1, 2, 2, 3, 2, 2, 2, 3, 2, 2, 2, 3, 4, 5, 6, 4, 4, 3, 2},
				predicate: func(i int) bool { return i > 2 },
			},
			want: []int{3, 3, 3, 4, 5, 6, 4, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.input, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		input      []int
		aggregator func(aggregate int, item int) int
		first      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple summation",
			args: args{
				input: []int{1, 2, 3, 4},
				aggregator: func(aggregate int, item int) int {
					return aggregate + item
				},
				first: 0,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.input, tt.args.aggregator, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple case",
			args: args{
				input: []int{1, 2, 3, 4, 1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
