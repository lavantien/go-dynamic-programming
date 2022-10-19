package main

/*
**Name**: Sum

**Description**: Given an integer N, find sum of 1 + 2 + 3 + ... + N, if N < 1 then the sum should be 0

**Examples**:

*Input*:
N = 10
*Output*:
55
*/
func Sum(n int) int {
	if n < 1 {
		return 0
	}

	acc := make([]int, n+1)
	acc[0] = 0
	for i := 1; i < len(acc); i++ {
		acc[i] = acc[i-1] + i
	}

	return acc[n]
}
