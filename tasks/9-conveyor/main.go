package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	/*Канал для записи данных*/
	dataCh := make(chan int)
	/*Канал для перемножения данных*/
	multiplierCh := make(chan int)

	go stage1(data, dataCh)
	go stage2(dataCh, multiplierCh)
	fmt.Print("Result:")

	for value := range multiplierCh {
		fmt.Print(" ", value)
	}
}

/*Пишет числа из слайса в канал*/
func stage1(data []int, dataCh chan int) {
	for _, value := range data {
		dataCh <- value
	}
	/*Закрываем канал с данными, чтобы stage2 не заблокировался*/
	defer close(dataCh)
}

/*Берет данные из канала dataCh и отправляет, помноженные на 2, данные в multiplierCh*/
func stage2(dataCh chan int, multiplierCh chan int) {
	for value := range dataCh {
		multiplierCh <- value * 2
	}
	/*Закрываем канал-множитель, чтобы main не заблокировался*/
	defer close(multiplierCh)
}
