package main

import "fmt"

func handleTemp(data []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, i := range data {
		//находим ключи (шаг 10)
		step := (int(i) / 10) * 10   //int(i) убирает все после точки, деление получает старший разряд. после умнож получаем -20 30 и тд
		m[step] = append(m[step], i) //step ключ (10 20 30 ...) , значение - массив температур float64
	}

	return m
}

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(handleTemp(data))
}
