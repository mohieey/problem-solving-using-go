package main

import (
	"strings"
)

func isValid(s string) bool {
	var left, right string = "{[(", "}])"
	lefts := []string{}

	for _, c := range s {
		if strings.Contains(left, string(c)) {
			lefts = append(lefts, string(c))
		}

		if strings.Contains(right, string(c)) {

			if len(lefts) == 0 {
				return false
			}

			if strings.Index(left, lefts[len(lefts)-1]) != strings.Index(right, string(c)) {
				return false
			}
			lefts = lefts[:len(lefts)-1]
		}

	}

	if len(lefts) == 0 {
		return false
	}

	return true
}
