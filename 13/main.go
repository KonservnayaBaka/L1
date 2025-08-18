package main

import "fmt"

func main() {
	a := 5
	b := 7

	a = a - b
	b = a + b
	a = b - a
	fmt.Println(a)
	fmt.Println(b)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	c := 4
	d := 8

	c = c ^ d
	d = c ^ d
	c = c ^ d
	fmt.Println(c)
	fmt.Println(d)
}
