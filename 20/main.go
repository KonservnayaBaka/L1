package main

import "fmt"

func reverseAll(s string) string {
	runeStr := []rune(s)

	for i := 0; i < len(runeStr)/2; i++ {
		runeStr[i], runeStr[len(runeStr)-1-i] = runeStr[len(runeStr)-1-i], runeStr[i]
	}

	runeStr = reverseWords(runeStr)

	return string(runeStr)
}

func reverseWords(s []rune) []rune {
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			end := i - 1
			for start < end {
				s[start], s[end] = s[end], s[start]
				start++
				end--
			}
			start = i + 1
		}
	}
	return s
}

func main() {
	fmt.Println(reverseAll("snow dog sun"))
}
