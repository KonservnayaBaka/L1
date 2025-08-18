package main

import (
	"fmt"
	"strings"
)

func uniqueString(s string) bool {
	lowerStr := strings.ToLower(s)
	m := make(map[rune]struct{}) //Можно сделать и через bool, но зато Struct ничего не весит

	for _, v := range lowerStr {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(uniqueString("abcd"))
	fmt.Println(uniqueString("abCdefAaf"))
	fmt.Println(uniqueString("aabcd"))
}

//Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
//Удобно map.
//Если возьмем массив/слайс, то придется каждый раз перебирать его значения в поисках определенного символа. Скорость такого алгоритма O(n)
//Если мы берем map, то скорость алгоритма (в среднем) увеличивается до O(1), так как map основан на хеш таблице
