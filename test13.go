package main

import "fmt"

func someAction(v []int, b int) {
	fmt.Println("act v =", &v[0])
	v[0] = 100
	v = append(v, b)
	v[1] = 250
	fmt.Println("act v append =", &v[0])
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println("act a =", &a[0])
	someAction(a, 6)
	fmt.Println(a)
}
