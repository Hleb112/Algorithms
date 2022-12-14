package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{11, 12, 13, 14, 15, 16, 1, 2, 3, 4}
	set := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, v := range arr1 {
		set[v] = true
	}

	for _, v := range arr2 {
		if _, ok := set[v]; ok {
			set2[v] = true
		}
	}

	fmt.Println(set2)
}
