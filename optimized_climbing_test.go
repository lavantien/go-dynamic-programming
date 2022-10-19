package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOptimizedClimbing(t *testing.T) {
	type args struct {
		n int
		k int
		f []int
	}
	type want struct {
		res int
		pos []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "OK",
			args: args{
				n: 3,
				k: 2,
				f: []int{3, 2, 4},
			},
			want: want{
				res: 6,
				pos: []int{0, 2, 3},
			},
		},
		{
			name: "OK1",
			args: args{
				n: 8,
				k: 2,
				f: []int{3, 2, 4, 6, 1, 1, 5, 3},
			},
			want: want{
				res: 11,
				pos: []int{0, 2, 3, 5, 6, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, gotPos := OptimizedClimbing(tt.args.n, tt.args.k, tt.args.f); got != tt.want.res || !cmp.Equal(gotPos, tt.want.pos) {
				t.Errorf("OptimizedClimbing() = %v, %v, want %v, %v", got, gotPos, tt.want.res, tt.want.pos)
			}
		})
	}
}
