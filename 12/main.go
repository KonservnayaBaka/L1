package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := map[string]struct{}{} //Можно сделать и через bool, но зато Struct ничего не весит
	arr2 := make([]string, 0, len(arr))

	for _, v := range arr {
		if _, ok := m[v]; !ok {
			arr2 = append(arr2, v)
			m[v] = struct{}{}
		}
	}

	fmt.Println(arr2)
}
