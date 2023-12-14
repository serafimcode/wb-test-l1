package main

import "fmt"

var justString string

/*
Проблема изначального варианта была в том, что что v остается в памяти даже после выполнения ф-ии.
Так происходит из-за устройства слайсов. Слайса по сути является структурой c указателем не массив.

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

Поэтому, когда в justString записывался слайс v[:100], мы все лишь создаем указатель на результат createHugeString.
*/

func someFunc() {
	/*Сделаем так, чтобы сборщик муссора почистил эту строку после выполнения ф-ии*/
	v := createHugeString(1 << 10)
	/*Создаем слайс байт(которым на самом деле является строка) необходимой длины*/
	c := make([]byte, len(v[0:100]))
	/*Из строки v скопируем элементы в c*/
	copy(c, v)
	/*Теперь в с у нас лежит на сегмент ссылающийся на v, а новый массив, с теми же элементами,
	что и в v(первые 100 элементов)*/
	justString = string(c)
}

func createHugeString(len int) string {
	str := "s"
	for i := 0; i < len; i++ {
		str += "s"
	}
	return str
}

func main() {
	someFunc()
	fmt.Println(justString)
}
