package main

import (
	"fmt"
)

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} //-27.0, -25.4, -21.0, 13.0, 15.5, 19.0, 24.5, 32.5
	m := map[int][]float32{}

	//slices.Sort(arr)

	for i := 0; i < len(arr); i++ {
		vInt := int(arr[i]/10) * 10
		m[vInt] = append(m[vInt], arr[i])
	}

	fmt.Println(m)
}
