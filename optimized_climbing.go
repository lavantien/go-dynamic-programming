package main

/*
**Name**: Optimized Climbing

**Description**: Given an integer N as the number of staircases, supposed you can only walk either one step at a time or two steps at a time, or up to k step at a time. And there is certain cost associate with the stairs, given in the integer array F. Calculate the route with minimum cost to reach the top from the bottom of the stair. If N < 0, you can assume that the result is 0. And there is 0 cost if you are on the ground

**Examples**:

*Input 1*:
N = 3, K = 2, F = [3, 2, 4]
*Output 1*:
6
[0, 2, 3]

*Input 1*:
N = 8, K = 2, F = [3, 2, 4, 6, 1, 1, 5 ,3]
*Output 1*:
11
[0, 2, 3, 5, 6, 8]
*/
// return results and route
func OptimizedClimbing(n int, k int, f []int) (int, []int) {
	/*
		Framework:
		1. Objective function:
			f(i) is the minimum cost at i
		2. Base cases:
			f(0) = 0
			f(1) = F(1)
		3. Recurrence relation for optimized objective function:
			f(i) = F(i) + min(F(i-k..i-1))
			route[i] = i-k+minPos
			route[n] = n
		4. Order of execution:
			bottom-up
		5. Answer position
			f(n), route
	*/
	if n < 0 || k < 1 {
		return 0, nil
	}
	if n == 0 {
		return 0, nil
	}

	f = append([]int{0}, f...)
	acc := make([]int, n+1)
	indexes := []int{}
	acc[0] = 0
	acc[1] = f[1]
	for i := 2; i < len(acc); i++ {
		left := i - k
		if left < 0 {
			left = 0
		}
		min, index := minOfVec(acc[left:i])
		index = left + index
		if !existed(indexes, index) {
			indexes = append(indexes, index)
		}
		acc[i] = f[i] + min
	}
	indexes = append(indexes, n)

	return acc[n], indexes
}

// return min and its position
func minOfVec(vec []int) (int, int) {
	min := vec[0]
	index := 0

	for i, ele := range vec {
		if ele < min {
			min = ele
			index = i
		}
	}

	return min, index
}

func existed(vec []int, x int) bool {
	for _, ele := range vec {
		if ele == x {
			return true
		}
	}

	return false
}
