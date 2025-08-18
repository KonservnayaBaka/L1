package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var i int
	fmt.Scan(&i)

	fmt.Println("Начальный массив: ", arr)

	arr = append(arr[:i], arr[i+1:]...) //Утечки не будет если брать ссылки на удаленные элементы
	fmt.Println("append: ", arr)
	fmt.Printf("len=%d, cap=%d\n", len(arr), cap(arr))

	copy(arr2[i:], arr2[i+1:])
	arr2[len(arr)-1] = 0
	arr2 = arr2[:len(arr2)-1]
	fmt.Println("copy: ", arr2)
	fmt.Printf("len=%d, cap=%d\n", len(arr2), cap(arr2))
}
