package main

import (
	"fmt"
)

func bfs(nums [][]int) {
	vals := []int{}
	seen := make([][]bool, len(nums))
	queue := [][]int{{0, 0}}

	for row, _ := range seen {
		seen[row] = make([]bool, len(nums[0]))
	}

	for len(queue) > 0 {
		position := queue[0]
		queue = queue[1:]

		row := position[0]
		col := position[1]

		if (row >= len(nums)) ||
			(row < 0) ||
			(col >= len(nums[0])) ||
			(col < 0) {
			continue
		}

		if seen[row][col] {
			continue
		}

		seen[row][col] = true
		vals = append(vals, nums[row][col])

		for _, direction := range directions {
			queue = append(queue, []int{row + direction[0], col + direction[1]})
		}

	}
	fmt.Println(vals)
}
