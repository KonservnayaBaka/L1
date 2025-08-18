package main

import "fmt"

func stringReverse(s string) string {
	runeStr := []rune(s)

	for i := 0; i < len(runeStr)/2; i++ {
		runeStr[i], runeStr[len(runeStr)-1-i] = runeStr[len(runeStr)-1-i], runeStr[i]
	}

	return string(runeStr)
}

func main() {
	fmt.Println(stringReverse("главрыба"))
}
