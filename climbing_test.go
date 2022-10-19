package main

import "testing"

func TestClimbing(t *testing.T) {
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
				n: 4,
			},
			want: 5,
		},
		{
			name: "OK1",
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			name: "Zero",
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "One",
			args: args{
				n: 1,
			},
			want: 1,
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
			if got := Climbing(tt.args.n); got != tt.want {
				t.Errorf("Climbing() = %v, want %v", got, tt.want)
			}
		})
	}
}
