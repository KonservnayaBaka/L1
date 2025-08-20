package main

import "strings"

var justString string //Глобальные переменные это плохо так как сильно вредит инкапсуляции, я бы переменную передавал в функцию

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100]	Строки ссылаются на массив байт. При таком представлении в justString будет записаны числа :100, но в буфере все равно останется ссылка на полный массив байтов
	//В итоге получается утечка памяти
	//Можно использовать такой подход для избежания такой ситуации
	justString = strings.Clone(v[:100])
	//Либо через copy
	byteSlice := make([]byte, 100)
	copy(byteSlice, v[:100])
	justString = string(byteSlice)
}

func main() {
	someFunc()
}
