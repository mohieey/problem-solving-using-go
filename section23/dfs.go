package main

import "fmt"

func dfs(nums [][]int) {
	vals := []int{}
	seen := make([][]bool, len(nums))
	for row, _ := range seen {
		seen[row] = make([]bool, len(nums[0]))
	}
	traverseDfs(0, 0, nums, seen, &vals)
	fmt.Println(vals)
}

func traverseDfs(row int, col int, nums [][]int, seen [][]bool, vals *[]int) {
	if (row >= len(nums)) ||
		(row < 0) ||
		(col >= len(nums[0])) ||
		(col < 0) {
		return
	}

	if seen[row][col] {
		return
	}
	*vals = append(*vals, nums[row][col])
	seen[row][col] = true
	for _, direction := range directions {

		traverseDfs(row+direction[0], col+direction[1], nums, seen, vals)

	}
}
