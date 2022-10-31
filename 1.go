package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
	Action
}

type Action struct {
}

func (a Action) walk() {
	fmt.Println("walking")
}

func main() {
	h := Human{
		name:   "Slava",
		age:    0,
		Action: Action{},
	}
	h.walk()
}
