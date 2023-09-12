package main

func minRemoveToMakeValid(s string) string {
	stack := []int{}
	i := 0
	for i < len(s) {
		char := string(s[i])
		if char == "(" {
			stack = append(stack, i)
		} else if char == ")" && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		} else if char == ")" {
			s = s[:i] + s[i+1:]
			continue
		}
		i++
	}
	i = len(stack) - 1
	for i >= 0 {
		println(i, s)
		s = s[:stack[i]] + s[stack[i]+1:]
		i--
	}
	return s
}
