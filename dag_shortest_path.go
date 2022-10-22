package main

import (
	"fmt"
	"math"
)

func DAG_ShortestPath(w [][]int) int {
	n := len(w)
	if n == 0 {
		return n
	}
	f := make([]int, n)
	for i := range f {
		f[i] = math.MaxInt32
	}
	f[n-1] = 0
	from := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(w[i]); j++ {
			weight := w[i][j]
			if weight > 0 {
				if f[i] > weight+f[j] {
					f[i] = weight + f[j]
					from[i] = j
				}
			}
		}
	}

	fmt.Println(from)

	return f[0]
}
