package main

import "fmt"

func reverseString(str string) string {
	/*Строки в go являются []byte, поэтому поэтому сперва привем строку к []rune.
	rune представляет собой код из unicode, т.e. один символ. */
	runes := []rune(str)
	/*Теперь мы можем получить длину строки, а не слайсай байт*/
	l := len(runes)

	/*Идем в цикле, как бы с двух сторон, l-i-1 позволяет нам найти положение символа,
	который находится на расстоянии i от праваого края слайса*/
	for i := 0; i < l/2; i++ {
		runes[i], runes[l-i-1] = runes[l-i-1], runes[i]
	}

	/*кастим обратно в строку*/
	return string(runes)
}

func main() {
	s := "главрыба\u2318"

	fmt.Print(reverseString(s))
}
