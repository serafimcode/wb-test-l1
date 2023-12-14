package main

import (
	"fmt"
	"time"
)

func main() {
	/*Читаем из stdin количество секунд, которое будет работать программа*/
	fmt.Print("Seconds to live: ")
	var secondsToLive time.Duration
	fmt.Scanf("%d", &secondsToLive)

	/*Создаем таймер*/
	timer := time.NewTimer(secondsToLive * time.Second)
	/*Создаем канал с данными*/
	dataCh := make(chan string)

	fmt.Println("Started at ", time.Now())

	/*Go routine, который будет писать в канал*/
	go writeToChannel(dataCh)

	/*Создаем бесконечный цикл, в котором слушаем канал с данными канал таймера*/
	for {
		select {
		/*Если данные есть, выводим их stdout*/
		case data := <-dataCh:
			fmt.Println(data)
		/*Когда таймер отсчитает указанное при создании время,
		канал в таймере вернет временную метку*/
		case t := <-timer.C:
			fmt.Println("Timer stopped at", t)
			return
		}
	}
}

/*Ф-ия записи данных в канал взята из предыдущего задания*/
func writeToChannel(dataCh chan<- string) {
	for {
		time.Sleep(time.Second)
		dataCh <- fmt.Sprintf("%v", time.Now())
	}
}
