package main

import "fmt"

// Input: m x n grid; Output: number of paths from (0,0) to (m,n); Restriction: only downward and forward are allowed
func GridPaths(m int, n int) int {
	// function: res[i,j] = res[i,j-1] + res[i-1,j]
	// wall: res[i,j] = 1 if i == -1 || j == -1
	// result: res[m-1,n-1]
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	// setup wall
	for i := 0; i < m; i++ {
		res[i][0] = 1
	}
	for j := 0; j < n; j++ {
		res[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i][j-1] + res[i-1][j]
		}
	}

	printGrid(res)

	return res[m-1][n-1]
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j], ", ")
		}
		fmt.Println()
	}
}

// Input: m x n grid with weights; Output: Path with maximum weight; Restriction: only downward and forward are allowed
func MaxWeightedGridPath(grid [][]int) (int, [][]int) {
	// function: res[i, j] = grid[i, j] + max(res[i, j-1], res[i-1, j]); with i = 1 -> m-1 and j = 1 -> n-1
	// result: res[m-1, n-1]
	// wall: res[i, j] = 0 if i < 0 || j < 0
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = grid[i][j]
			if i-1 < 0 && j-1 < 0 {
				res[i][j] += 0
			} else if i-1 < 0 {
				res[i][j] += res[i][j-1]
			} else if j-1 < 0 {
				res[i][j] += res[i-1][j]
			} else {
				max := max(res[i-1][j], res[i][j-1])
				res[i][j] += max
			}
		}
	}

	return res[m-1][n-1], getPath(res, m-1, n-1, [][]int{})
}

func getPath(res [][]int, i int, j int, path [][]int) [][]int {
	if i == 0 && j == 0 {
		return append(path, []int{0, 0})
	}

	if i == 0 {
		path = getPath(res, i, j-1, path)
	} else if j == 0 {
		path = getPath(res, i-1, j, path)
	} else {
		if res[i-1][j] > res[i][j-1] {
			path = getPath(res, i-1, j, path)
		} else {
			path = getPath(res, i, j-1, path)
		}
	}

	return append(path, []int{i, j})
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func reverse[T any](vec []T) []T {
	if len(vec) == 0 {
		return []T{}
	}

	for i, j := 0, len(vec)-1; i < j; i, j = i+1, j-1 {
		vec[i], vec[j] = vec[j], vec[i]
	}

	return vec
}
