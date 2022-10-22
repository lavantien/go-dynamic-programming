package main

import "math"

// returns the minimum number of coins to make the change of size n
func ChangeMaking(coins []int, n int) int {
	f := make([]int, n+1)
	f[0] = 0
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt64
		for _, coin := range coins {
			if i-coin < 0 || 1+f[i-coin] == math.MaxInt64 {
				continue
			}
			if i >= coin {
				f[i] = int(math.Min(float64(f[i]), float64(1+f[i-coin])))
			}
		}
	}
	if f[n] == math.MaxInt64 {
		return -1
	}
	return f[n]
}
