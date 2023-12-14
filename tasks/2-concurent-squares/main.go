package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	out := make(chan int)

	/*В цикле запускаем go routine на каждый элемент nums*/
	for _, num := range nums {
		go square(num, out)
	}

	/*В цикле получаем значения из канала*/
	for i := 0; i < len(nums); i++ {
		/* Оператор <- блокирует выполнение главной go routine пока не придет значение из канала out*/
		fmt.Println(<-out)
	}
}

/*
Ф-ия квадрата принимает число для возведения в квадрат и канал,
в который будет отправлен результат. Канал типизирован, как канал принимающий int зн-я,
т.е. внутри ф-ии нельзя получить значения из этого канала
*/
func square(num int, out chan<- int) {
	out <- num * num
}
