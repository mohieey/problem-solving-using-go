package main

import "fmt"

func main() {
	fmt.Printf("numOfMinutes: %v\n", numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	list := adjacencyListBuilder(manager)
	return dfs(list, headID, informTime)

}

func adjacencyListBuilder(manager []int) [][]int {
	list := make([][]int, len(manager))
	for i, m := range manager {
		if m == -1 {
			continue
		}
		list[m] = append(list[m], i)
	}
	return list
}

func dfs(list [][]int, employee int, informTime []int) int {
	local := 0
	for _, managedEmp := range list[employee] {
		t := dfs(list, managedEmp, informTime)
		if t > local {
			local = t
		}
	}

	return local + informTime[employee]
}
