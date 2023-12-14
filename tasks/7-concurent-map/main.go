package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	myMap := map[int]int{}
	/*Создаем мьютекс, что 2 горутина не писали в одну область памяти одновременно*/
	mutex := sync.Mutex{}
	/*WaitGroup чтобы дождаться завершения работы всех горутин*/
	wg := sync.WaitGroup{}

	/*Создаем 10 горутин, добавляя их в wg*/
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go updateMap(myMap, &mutex, &wg, i)
	}
	wg.Wait()
	fmt.Printf("myMap: %+v", myMap)
}

func updateMap(m map[int]int, mutex *sync.Mutex, wg *sync.WaitGroup, id int) {
	/*Создаем рандомные ключ и зн-я для нашей мапы*/
	key := rand.Intn(10)
	value := rand.Intn(100)

	/*Захватываем мьютекс*/
	mutex.Lock()

	/*По завершению горутины отпускаем мьютекс и уменьшаем счетчик wg*/
	defer func() {
		wg.Done()
		mutex.Unlock()
	}()

	m[key] = value
	fmt.Printf("Worker %v set map[%v] to %v\n", id, key, value)
}
