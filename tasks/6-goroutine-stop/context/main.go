package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	/*Создаем контекст с таймаутом 5 секунд,
	т.е. канал контекста ctx.Done() сообщит нам о завершении через 5 секунд*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	/*Создаем WaitGroup, чтобы дождаться конца работы goroutine*/
	wg := sync.WaitGroup{}
	go worker(ctx, &wg)
	wg.Add(1)

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	/*Крутим бесконечный цикл*/
	for {
		/*Селект ждет, пока один из каналов не отдаст зн-е*/
		select {
		/*Когда пройдет таймаут контекста, канал закроется и выполнится */
		case <-ctx.Done():
			fmt.Println("Finishing goroutine")
			wg.Done()
			return
		/*Выполнится, если нет готовых каналов*/
		default:
			fmt.Println("Sill works")
			time.Sleep(1 * time.Second)
		}
	}
}
