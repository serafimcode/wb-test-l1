package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	/*Разбиваем строку на подстроки, разделенные пробелом*/
	res := strings.Fields(str)
	l := len(res)

	/*Повторям трюк из прошлого задания*/
	for i := 0; i < l/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}

	/*Склеиваем строки в одну строку*/
	return strings.Join(res, " ")
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverseString(str))
}
