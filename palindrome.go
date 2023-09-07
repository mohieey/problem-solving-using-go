package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	pattern := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = pattern.ReplaceAllString(strings.ToLower(s), "")

	pL, pR := 0, (len(s) - 1)

	for pL < pR {

		if s[pL] == s[pR] {
			pL++
			pR--

		} else {
			return false
		}
	}
	return true
}
