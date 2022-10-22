package main

import (
	"testing"
)

func TestDAG_ShortestPath(t *testing.T) {
	type args struct {
		w [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				w: [][]int{
					{0, 2, 3, 0, 0, 0},
					{0, 0, 0, 5, 3, 0},
					{0, 1, 0, 4, 11, 0},
					{0, 0, 0, 0, 0, 4},
					{0, 0, 0, 3, 0, 7},
					{0, 0, 0, 0, 0, 0},
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DAG_ShortestPath(tt.args.w); got != tt.want {
				t.Errorf("DAG_ShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
