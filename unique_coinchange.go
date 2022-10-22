package main

// returns the unique number of ways to make a change of size n
func UniqueCoinChange(coins []int, n int) int {
	f := make([]int, n+1)
	f[0] = 1

	for j := 0; j < len(coins); j++ {
		for i := 1; i <= n; i++ {
			if i >= coins[j] {
				f[i] += f[i-coins[j]]
			}
		}
	}

	return f[n]
}
