package main

import "fmt"

func main() {

	graph := [][]int{
		{1, 3},
		{0},
		{3, 8},
		{0, 2, 4, 5},
		{3, 6},
		{3},
		{4, 7},
		{6},
		{2},
	}

	traverse(graph)

}

func traverse(graph [][]int) {
	values := []int{}
	seen := make([]bool, len(graph))

	dfs(graph, 0, seen, &values)
	fmt.Printf("values: %v\n", values)
}

func dfs(graph [][]int, vertex int, seen []bool, values *[]int) {
	if seen[vertex] {
		return
	}
	*values = append(*values, vertex)
	seen[vertex] = true

	for _, nextVertex := range graph[vertex] {
		dfs(graph, nextVertex, seen, values)
	}
}
