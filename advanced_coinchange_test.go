package main

import "testing"

func TestAdvancedCoinChange(t *testing.T) {
	type args struct {
		coins []int
		n     int
		s     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				coins: []int{1, 2, 3, 5},
				n:     7,
				s:     3,
			},
			want: 9,
		},
		{
			name: "Edge",
			args: args{
				coins: []int{1, 2, 3, 5},
				n:     0,
				s:     0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdvancedCoinChange(tt.args.coins, tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("AdvancedCoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
