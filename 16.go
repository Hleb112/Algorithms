package main

import (
	"fmt"
	"math/rand"
	"time"
)

// главная сложность - выбор опорного элемента
// худший случай n*n
// средний n log2 n
func qs(data []int) []int {
	// базовый случай
	if len(data) < 2 {
		return data
	}

	less := make([]int, 0)
	greater := make([]int, 0)

	// получаем худший случай
	pivot := data[0]
	for _, i := range data[1:] {
		if i > pivot {
			// больше опорного
			greater = append(greater, i)
		} else {
			//меньших опорного
			less = append(less, i)
		}
	}
	// сортируем меньшие и большие
	data = append(qs(less), pivot)
	data = append(data, qs(greater)...)
	return data
}

func qSort(data []int) []int {
	// базовый случай
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1

	// рекурсивый случай
	rand.Seed(time.Now().UnixNano())
	pivot := rand.Int() % len(data)

	// двигаем вправо опорный элемент
	data[pivot], data[right] = data[right], data[pivot]

	// двагаем элементы меньше опорного влево
	for i := range data {
		if data[i] < data[right] {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}

	// left - позиция последнего элемента, который больше опорного, установим в неё опорный элемент
	data[left], data[right] = data[right], data[left]

	qSort(data[:left])
	qSort(data[left+1:])

	return data
}

func main() {
	data := []int{2, 4, 3, 8, 5, 4, 6, 1, 7}

	//sort.Ints(data)
	fmt.Println(qs(data))
}
