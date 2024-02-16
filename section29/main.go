package main

import "fmt"

func main() {
	p := [][]int{
		{1, 0},
		{2, 1},
		{2, 5},
		{0, 3},
		{4, 3},
		{3, 5},
		{4, 5},
	}

	fmt.Printf("canFinish(6, p): %v\n", canFinish(6, p))

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDeg, Adj := toInDegreeAndAdj(numCourses, prerequisites)
	zeros := []int{}
	for vertex, deg := range inDeg {
		if deg == 0 {
			zeros = append(zeros, vertex)
		}
	}

	processedVertices := 0

	for len(zeros) > 0 {
		vertex := zeros[0]
		zeros = zeros[1:]
		processedVertices++

		connected := Adj[vertex]
		for _, connectedVer := range connected {
			inDeg[connectedVer]--
			if inDeg[connectedVer] == 0 {
				zeros = append(zeros, connectedVer)
			}
		}

	}

	return processedVertices == numCourses
}

func toInDegreeAndAdj(numCourses int, prerequisites [][]int) ([]int, [][]int) {
	inDegree := make([]int, numCourses)
	Adj := make([][]int, numCourses)

	for _, courseAndPre := range prerequisites {
		course := courseAndPre[0]
		prerequisite := courseAndPre[1]

		inDegree[course]++

		Adj[prerequisite] = append(Adj[prerequisite], course)
	}

	return inDegree, Adj
}
