package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) // Новое множество
	for i := 0; i < len(str); i++ {
		set[str[i]] = true
	}

	fmt.Println(set)

}
