package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(BinarySearch(nums, 0, len(nums)-1, 3))

	nums2 := []int{1, 2, 3, 3, 3, 3, 4, 5, 6, 7}

	s, e := SearchRange(nums2, 3)
	fmt.Println(s, e)

}

func SearchRange(nums []int, target int) (int, int) {
	firstPosition := BinarySearch(nums, 0, len(nums)-1, target)
	if firstPosition == -1 {
		return -1, -1
	}

	start, end, tempStart, tempEnd := firstPosition, firstPosition, 0, 0

	for start != -1 {
		tempStart = start
		start = BinarySearch(nums, 0, start-1, target)
	}
	start = tempStart

	for end != -1 {
		tempEnd = end
		end = BinarySearch(nums, end+1, len(nums)-1, target)
	}
	end = tempEnd

	return start, end
}

func BinarySearch(nums []int, left, right, target int) int {

	for left <= right {
		mid := int(math.Floor((float64(left+right) / 2)))
		val := nums[mid]
		switch {
		case val == target:
			return mid
		case target < val:
			right = mid - 1
		case target > val:
			left = mid + 1
		}
	}

	return -1
}
