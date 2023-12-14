package main

import (
	"fmt"
	"sync"
)

/*
Создаем структуру счетчика с RW мьютексом,
чтобы блокировать доступ при записи,
но не трогать при чтении
*/
type counter struct {
	value int
	mu    sync.RWMutex
}

/*Перед каждой операцией запии захватываем мьютекс,
после выполнения отпускаем*/

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

/*Простой воркер который будет изменять зн-е счетчика*/
func worker(wg *sync.WaitGroup, c *counter, id int) {
	defer wg.Done()

	if id > 3 {
		c.inc()
	} else {
		c.dec()
	}
}

func main() {
	c := counter{value: 5}
	wg := sync.WaitGroup{}

	/*Запускаем 5 горутин*/
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, &c, i)
	}

	/*Ждем завершения всех горутин*/
	wg.Wait()
	fmt.Println(c.value)
}
