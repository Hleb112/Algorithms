package main

import (
	"fmt"
)

func update(p *int) {
	fmt.Println("update 0 p = ", p)
	b := 2
	fmt.Println("update b =", &b)
	p = &b
	fmt.Println("update p = ", p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println("main p", p)
	update(p)
	fmt.Println("main p", p)
}
