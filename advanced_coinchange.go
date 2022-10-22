package main

// returns the total number of ways to make a change of size n, by using exactly s coins
func AdvancedCoinChange(coins []int, n int, s int) int {
	// f(i, k) = SumJ f(i-coins[j], k-1)
	// base: f(0, 0) = 1, f(i, 0) = 0
	f := make([][]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, s+1)
	}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= s; j++ {
			if i > 0 && j == 0 {
				f[i][j] = 0
				continue
			}
			for _, coin := range coins {
				if i-coin >= 0 {
					f[i][j] += f[i-coin][j-1]
				}
			}
		}
	}
	return f[n][s]
}
