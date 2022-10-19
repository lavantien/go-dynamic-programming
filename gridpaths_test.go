package main

import (
	"reflect"
	"testing"
)

func TestGridPaths(t *testing.T) {
	type args struct {
		m int
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
				m: 3,
				n: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GridPaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("GridPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxWeightedGridPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 [][]int
	}{
		{
			name: "OK",
			args: args{
				grid: [][]int{
					{0, 2, 2, 1},
					{3, 1, 1, 1},
					{4, 4, 2, 0},
				},
			},
			want: 13,
			want1: [][]int{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
				{2, 3},
			},
		},
		{
			name: "OK1",
			args: args{
				grid: [][]int{
					{0, 2, 2, 50},
					{3, 1, 1, 100},
					{4, 4, 2, 0},
				},
			},
			want: 154,
			want1: [][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 3},
				{2, 3},
			},
		},
		{
			name: "OK2",
			args: args{
				grid: [][]int{
					{0, 3, 2, 2},
					{2, 3, 3, 3},
					{2, 2, 2, 0},
				},
			},
			want: 12,
			want1: [][]int{
				{0, 0},
				{0, 1},
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MaxWeightedGridPath(tt.args.grid)
			if got != tt.want {
				t.Errorf("MaxWeightedGridPath() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MaxWeightedGridPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
