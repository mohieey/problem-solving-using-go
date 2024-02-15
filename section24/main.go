package main

import "fmt"

func main() {

	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Printf("numIslands(grid): %v\n", numIslands(grid))

}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func numIslands(grid [][]byte) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				count++
				bfs(row, col, &grid)
			}
		}
	}

	return count
}

func bfs(row int, col int, nums *[][]byte) {
	queue := [][]int{{row, col}}

	for len(queue) > 0 {
		position := queue[0]
		queue = queue[1:]

		row := position[0]
		col := position[1]

		if (row >= len(*nums)) ||
			(row < 0) ||
			(col >= len((*nums)[0])) ||
			(col < 0) {
			continue
		}

		if (*nums)[row][col] == '0' {
			continue
		}

		(*nums)[row][col] = '0'

		for _, direction := range directions {
			queue = append(queue, []int{row + direction[0], col + direction[1]})
		}

	}
}
