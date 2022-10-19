package main

import "testing"

func TestAdvancedClimbing(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				n: 5,
				k: 3,
			},
			want: 13,
		},
		{
			name: "OK1",
			args: args{
				n: 1000000,
				k: 2,
			},
			want: 2756670985995446685,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdvancedClimbing(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("AdvancedClimbing() = %v, want %v", got, tt.want)
			}
		})
	}
}
