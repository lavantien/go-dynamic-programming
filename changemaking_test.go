package main

import "testing"

func TestChangeMaking(t *testing.T) {
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
				n: 29,
				coins: []int{1,3,5},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeMaking(tt.args.coins, tt.args.n); got != tt.want {
				t.Errorf("ChangeMaking() = %v, want %v", got, tt.want)
			}
		})
	}
}
