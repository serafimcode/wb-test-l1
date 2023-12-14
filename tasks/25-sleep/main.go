package main

import (
	"fmt"
	"time"
)

/*
Крутим бесконечный цикл и на каждой итерации проверяем условие:
Время вызова + длительность сна - текущее время <= 0
*/
func sleep(d time.Duration) {
	start := time.Now()
	for {
		if start.Add(d).Sub(time.Now()) <= 0 {
			return
		}
	}
}

func main() {
	fmt.Println("Start: ", time.Now())
	sleep(3 * time.Second)
	fmt.Println("End: ", time.Now())
}
