package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2

	/*Складываем 2 числа и храним на месте одно из исходных*/
	num1 += num2
	/*В num1 лежит сумма num1'(предыдущее) и num2.
	Тогда, если мы отнимем num2, от num1, мы получим num1'*/
	num2 = num1 - num2
	/*Теперь в num2 лежит num1', num1 все еще num1'+num2.
	Вычтем num2 из num1, т.е. (num1' + num2) - num1'*/
	num1 -= num2

	fmt.Println("num1: ", num1)
	fmt.Println("num2: ", num2)
}
