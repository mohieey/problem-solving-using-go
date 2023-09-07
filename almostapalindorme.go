package main

import (
	"regexp"
	"strings"
)

func validPalindrome(s string) bool {
	pattern := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = pattern.ReplaceAllString(strings.ToLower(s), "")

	pL, pR := 0, (len(s) - 1)

	for pL < pR {

		if s[pL] == s[pR] {
			pL++
			pR--
		} else {
			return isSubPalindrome(s, pL+1, pR) || isSubPalindrome(s, pL, pR-1)
		}
	}
	return true

}

func isSubPalindrome(s string, pL int, pR int) bool {

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
