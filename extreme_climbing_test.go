package main

import "testing"

func TestExtremeClimbing(t *testing.T) {
	type args struct {
		n int
		k int
		f []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				n: 7,
				k: 3,
				f: []bool{false, true, false, true, true, false, false},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtremeClimbing(tt.args.n, tt.args.k, tt.args.f); got != tt.want {
				t.Errorf("ExtremeClimbing() = %v, want %v", got, tt.want)
			}
		})
	}
}
