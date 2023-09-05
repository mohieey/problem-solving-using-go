package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if len(s) < 2 {
		return 1
	}

	pL, pR := 0, 0

	seen := make(map[byte]int)
	maxLen := 0
	for pR < len(s) {
		p, ok := seen[s[pR]]
		if ok && p >= pL {
			pL = p + 1
			seen[s[pR]] = pR
		} else {
			seen[s[pR]] = pR
			maxLen = int(math.Max(float64(maxLen), float64(pR-pL+1)))
		}
		pR++

	}

	return maxLen
}
