package main

import "fmt"

func handleTemp(data []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, i := range data {
		//находим ключи (шаг 10)
		step := (int(i) / 10) * 10
		m[step] = append(m[step], i)
	}

	return m
}

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 9.9}

	fmt.Println(handleTemp(data))
}
