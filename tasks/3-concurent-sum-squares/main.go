package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	out := make(chan int)

	for _, num := range nums {
		go square(num, out)
	}
	/*Решение основано на предыдущем. Добавили переменную для суммы квадратов.*/
	var squareSum int

	for i := 0; i < len(nums); i++ {
		squareSum += <-out
	}
	fmt.Println(squareSum)
}

func square(num int, out chan<- int) {
	out <- num * num
}
