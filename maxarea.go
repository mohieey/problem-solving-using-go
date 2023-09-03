package main

import "math"

func maxArea(height []int) int {
	maxArea := 0

	pl, pr := 0, (len(height) - 1)

	for pl < pr {
		currentArea := int(math.Min(float64(height[pl]), float64(height[pr]))) * (pr - pl)
		maxArea = int(math.Max(float64(maxArea), float64(currentArea)))
		if height[pl] < height[pr] {
			pl++
		} else {
			pr--
		}
	}

	return maxArea
}
