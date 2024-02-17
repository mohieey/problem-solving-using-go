package main

import "math"

func main() {

}

func minCostClimbingStairsIT_OPT(cost []int) int {
	n := len(cost)
	v1 := cost[0]
	v2 := cost[1]
	for i := 2; i < n; i++ {
		val := cost[i] + int(math.Min(float64(v1), float64(v2)))
		v1 = v2
		v2 = val
	}

	return int(math.Min(float64(v1), float64(v2)))
}

func minCostClimbingStairsIT(cost []int) int {
	n := len(cost)
	mem := map[int]int{}
	for i := 0; i < n; i++ {
		if i < 2 {
			mem[i] = cost[i]
		} else {
			mem[i] = cost[i] + int(math.Min(float64(mem[i-1]), float64(mem[i-2])))
		}
	}

	return int(math.Min(float64(mem[n-1]), float64(mem[n-2])))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	mem := map[int]int{}
	return int(math.Min(float64(min(n-1, cost, mem)), float64(min(n-2, cost, mem))))
}

func min(i int, cost []int, mem map[int]int) int {
	if i < 0 {
		return 0
	}

	if i == 0 || i == 1 {
		return cost[i]
	}

	val, ok := mem[i]
	if ok {
		return val
	}

	val = cost[i] + int(math.Min(float64(min(i-1, cost, mem)), float64(min(i-2, cost, mem))))
	mem[i] = val
	return val
}
