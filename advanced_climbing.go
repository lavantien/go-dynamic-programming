package main

/*
**Name**: Advanced Climbing

**Description**: Given an integer N as the number of staircases, supposed you can only walk either one step at a time or two steps at a time, or up to k step at a time. Calculate the total number of ways to reach the top from the bottom of the stair. If N < 0, you can assume that the result is 0. And there is 1 way of climbing 0 staircase

**Examples**:

*Input 1*:
N = 5, K = 3
*Output 1*:
13

*Input 2*:
N = 1000000, K = 2
*Output 2*:
2756670985995446685
*/
func AdvancedClimbing(n int, k int) int {
	/*
		Framework:
		1. Objective function:
			f(i) is the number of distinct ways to reach the i-th stair by making 1 to k steps
		2. Base cases:
			f(0) = 1
			f(1) = 1
		3. Recurrence relation for optimized objective function:
			f(n) = f(n-1) + f(n-2) + ... + f(n-k)
		4. Order of execution:
			bottom-up
		5. Answer position
			f(n)
	*/
	if n < 0 || k < 1 {
		return 0
	}
	if n == 0 {
		return 1
	}

	acc := make([]int, n+1)
	acc[0] = 1
	acc[1] = 1
	for i := 2; i < len(acc); i++ {
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				continue
			}
			acc[i] += acc[i-j]
		}
	}

	return acc[n]
}
