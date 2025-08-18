package main

import "fmt"

func main() {
	m := map[string]interface{}{}
	ch := make(chan int, 1)
	ch <- 9

	m["one"] = 1
	m["two"] = "2"
	m["three"] = true
	m["four"] = 4.4
	m["five"] = ch

	for k, v := range m {
		switch val := v.(type) {
		case int:
			fmt.Printf("int type in key - %s:%d\n", k, val)
		case string:
			fmt.Printf("string type in key - %s:%s\n", k, val)
		case bool:
			fmt.Printf("bool type in key -  %s:%t\n", k, val)
		case float64:
			fmt.Printf("float type in key - %s:%.2f\n", k, val)
		case chan int:
			fmt.Printf("chan int type in key - %s:%v\n", k, <-val)
		case chan string:
			fmt.Printf("chan string type in key - %s:%v\n", k, <-val)
		case chan bool:
			fmt.Printf("chan bool type in key - %s:%v\n", k, <-val)

		}
	}

	close(ch)
}
