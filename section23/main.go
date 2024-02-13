package main

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	nums := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	dfs(nums)
	bfs(nums)

}
