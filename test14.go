package main

import "fmt"

type zxc func(a int)

func main() {
	slice := []string{"a", "a"}
	func(slice []string) {

		slice[0] = "b"
		slice = append(slice, "a")
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)

}
