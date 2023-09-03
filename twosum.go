package main

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	iteratedValues := make(map[int]int)

	for i, v := range nums {
		difference := target - v
		j, ok := iteratedValues[v]
		if ok {
			return []int{j, i}
		}
		iteratedValues[difference] = i
	}

	return []int{}
}
