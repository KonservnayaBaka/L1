package main

import "fmt"

func setBit(num int64, i int, val int) int64 {
	if i < 1 || i > 64 { //0...63
		return num
	}

	u := uint64(num)
	mask := uint64(1) << uint(i-1)

	if val == 1 {
		u |= mask
	} else if val == 0 {
		u &^= mask
	}
	return int64(u)
}

func main() {
	{
		num := int64(5)
		i := 1
		val := 0
		result := setBit(num, i, val)
		fmt.Printf("Было: %d: Стало: %d\n", num, result)
	}

	{
		num := int64(-5)
		i := 1
		val := 0
		result := setBit(num, i, val)
		fmt.Printf("Было: %d: Стало: %d\n", num, result)
	}
}
