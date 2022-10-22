package main

import "testing"

func TestUniqueCoinChange(t *testing.T) {
	type args struct {
		coins []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test case #1",
			args: args{
				n:     75,
				coins: []int{1, 2, 3, 5},
			},
			want: 2894,
		},
		{
			name: "simple test case #1",
			args: args{
				n:     75,
				coins: []int{2, 3, 5},
			},
			want: 107,
		},
		{
			name: "simple test case #2",
			args: args{
				n:     10,
				coins: []int{4, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueCoinChange(tt.args.coins, tt.args.n); got != tt.want {
				t.Errorf("UniqueCoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
