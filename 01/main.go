package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) InfoAboutHuman() {
	fmt.Printf("%s is %d years old\n", h.name, h.age)
}

type Action struct {
	Human
	Run  func()
	Jump func()
	Play func()
}

func main() {
	a := Action{
		Human: Human{
			name: "Teodor",
			age:  20,
		},
		Run:  func() { fmt.Println("Running...") },
		Jump: func() { fmt.Println("Jumping...") },
		Play: func() { fmt.Println("Playing...") },
	}

	a.InfoAboutHuman()
	a.Run()
	a.Jump()
	a.Play()
}
