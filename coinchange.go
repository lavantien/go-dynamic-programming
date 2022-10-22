package main

// returns the total number of ways to make a change of size n
func CoinChange(coins []int, n int) int {
	// f(i) is the total number of ways to make a change of size i
	// f[i] = SumJ f[i - coins[j]]
	// base: f[0] = 1
	f := make([]int, n+1)
	f[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				f[i] += f[i-coins[j]]
			}
		}
	}

	return f[n]
}
