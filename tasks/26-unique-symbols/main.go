package main

import (
	"fmt"
	"strings"
)

func checkUniqueness(str string) bool {
	/*Приводим строку к нижнему регистру*/
	str = strings.ToLower(str)

	/*Заводим мапу, ключ - символ(руна), зн-е - количество вхождений символа*/
	uniqueMap := map[rune]int{}

	/*Идем по страке через range. Range перебирает строку по рунам.
	Заполняем мапу, счетчик символа равен 0, ставим 1, иначе инкрементим счетчик*/
	for _, ch := range str {
		if uniqueMap[ch] == 0 {
			uniqueMap[ch] = 1
		} else {
			uniqueMap[ch]++
		}
	}

	/*Идем по мапе, первое вхождение со значением больше 1 - был повтор*/
	for _, v := range uniqueMap {
		if v > 1 {
			return false
		}
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(checkUniqueness(str1))
	fmt.Println(checkUniqueness(str2))
	fmt.Println(checkUniqueness(str3))
}
