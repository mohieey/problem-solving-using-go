package main

func trap(heights []int) int {

	totalWater, maxLeft, maxRight, left, right := 0, 0, 0, 0, (len(heights) - 1)

	for left < right {
		if heights[left] < heights[right] {
			if heights[left] > maxLeft {
				maxLeft = heights[left]
			} else {
				totalWater += maxLeft - heights[left]
			}
			left++
		} else {
			if heights[right] > maxRight {
				maxRight = heights[right]
			} else {
				totalWater += maxRight - heights[right]
			}
			right--
		}
	}

	return totalWater
}
