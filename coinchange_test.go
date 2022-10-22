package main

import "testing"

func TestCoinChange(t *testing.T) {
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
			name: "OK",
			args: args{
				coins: []int{1, 3, 5, 10},
				n:     3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange(tt.args.coins, tt.args.n); got != tt.want {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
