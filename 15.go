var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}

func main() {
someFunc()
}
*/

/*
	1) могут быть проблемы со строкой (не utf). Лучше кастануть в []rune
	2) из-за глобальной переменной возникнут проблемы со сборкой мусора
	3) justString = v[:100] - получит слайс, который указывает на большую строку -> в памяти будет храниться
	целый массив созданный createHugeString, а не только интересующая нас часть в 100 элементов
	4) скорее всего все будет располагаться на хипе

	исправить можно:
	1) сделать justString локальной
	// закнем нужные нам элементы в слайс
	2) out := append([]byte{}, v[:100]...)
       return out
	// если решим оставить глобальную переменную
	3) out := append([]byte{}, v[:100]...)
       justString = out
*/