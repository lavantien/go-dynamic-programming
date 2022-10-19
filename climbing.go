package main

/*
**Name**: Climbing

**Description**: Given an integer N as the number of staircases, supposed you can only walk either one step at a time or two steps at a time. Calculate the total number of ways to reach the top from the bottom of the stair. If N < 0, you can assume that the result is 0. And there is 1 way to climb a 0 starcase

**Examples**:

*Input 1*:
N = 2
*Output 1*:
2

*Input 2*:
N = 0
*Output 2*:
1

*Input 2*:
N = 4
*Output 2*:
5
*/
func Climbing(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	acc := make([]int, n+1)
	acc[0] = 1
	acc[1] = 1
	for i := 2; i < len(acc); i++ {
		acc[i] = acc[i-1] + acc[i-2]
	}

	return acc[n]
}
