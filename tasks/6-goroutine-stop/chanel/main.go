package main

import (
	"fmt"
	"time"
)

func main() {
	/*Создаем канал, который просигналит goroutine, когда нужно остановиться*/
	stopCh := make(chan bool)
	/*Запускаем goroutine и передаем ему сигнал*/
	go worker(stopCh)
	/*Даем goroutine немного поработать*/
	time.Sleep(5 * time.Second)
	/*Передаем сигнал об остановке, что передадим- не важно, важен сам факт того, что сигнал сработал*/
	//stopCh <- true // можно заменить на close(stopCh), тогда канал закроется и
	close(stopCh)
}

func worker(stopCh <-chan bool) {
	/*Крутим бесконечный цикл*/
	for {
		/*Селект ждет, пока один из каналов не отдаст зн-е*/
		select {
		/*Если придет зн-е из stopCh, мы выйдем из цикла*/
		case <-stopCh:
			fmt.Println("Finishing goroutine")
			return
		/*Выполнится, если нет готовых каналов*/
		default:
			fmt.Println("Sill works")
			time.Sleep(1 * time.Second)
		}
	}
}
