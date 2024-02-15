package main

import "fmt"

func main() {
	// grid := [][]int{
	// 	{2, 1, 1},
	// 	{1, 1, 0},
	// 	{0, 1, 1},
	// }

	grid := [][]int{
		{1, 2},
	}
	fmt.Println(orangesRotting(grid))
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func orangesRotting(grid [][]int) int {
	minutes := 0
	fresh := 0
	rotten := [][]int{}
	rottenLen := len(rotten)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			switch grid[row][col] {
			case 1:
				fresh++
			case 2:
				rotten = append(rotten, []int{row, col})
				rottenLen = len(rotten)
			}
		}
	}

	if rottenLen == 1 && fresh == 0 {
		return 0
	}

	for len(rotten) > 0 {
		position := rotten[0]
		rotten = rotten[1:]
		rottenLen = len(rotten)

		if rottenLen == 0 {
			minutes++
		}

		row := position[0]
		col := position[1]

		for _, direction := range directions {
			if (row+direction[0] >= len(grid)) ||
				(row+direction[0] < 0) ||
				(col+direction[1] >= len(grid[0])) ||
				(col+direction[1] < 0) {
				continue
			}

			if grid[row+direction[0]][col+direction[1]] != 1 {
				continue
			}

			grid[row+direction[0]][col+direction[1]] = 2
			fresh--

			rotten = append(rotten, []int{row + direction[0], col + direction[1]})
			rottenLen = len(rotten)
		}

	}

	if fresh != 0 {
		return -1
	}

	return minutes

}
