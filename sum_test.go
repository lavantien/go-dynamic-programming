package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				n: 10,
			},
			want: 55,
		},
		{
			name: "One",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "Zero",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "Negative",
			args: args{
				n: -1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.n); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
