package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	/*Канал, в который будут записываться данные и из которого воркеры будут эти данные брать*/
	dataChannel := make(chan string)
	/*Завершение воркеров реализуем через WaitGroup,
	так как этот механизм специально прдназначен для работы
	с коллекцией goroutine*/
	var wg sync.WaitGroup

	/*Читаем из stdin количество воркеров*/
	fmt.Print("Chose number of workers: ")
	var numOfWorkers int
	fmt.Scanf("%d", &numOfWorkers)
	fmt.Println(numOfWorkers)

	/*Запускаем наши воркеры, инкрементируя счетчик wg*/
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	/*Запускаем goroutine который будет писать данные в канал*/
	go writeToChannel(dataChannel)

	/*Запускаем слушатель SIGINT*/
	listenInterrupt(dataChannel)

	/*Ждем пока не завершатся все goroutine*/
	wg.Wait()
	fmt.Println("All workers have finished. Exiting...")
}

/*Наш воркер, принимает id, канал для чтения данных и указатель на WaitGroup*/
func worker(workerId int, dataCh <-chan string, wg *sync.WaitGroup) {
	/*Когда dataCh будет закрыт, мы выйдем из цикла и ф-ия worker завершится.
	В здесь мы уменьшаем счетчик WaitGroup и логируем завершение воркера*/
	defer func() {
		fmt.Printf("Worker %d finished:\n", workerId)
		wg.Done()
	}()

	/*Получаем данные из канала данных, пока канал не закроется*/
	for data := range dataCh {
		fmt.Printf("Worker %d received: %s\n", workerId, data)
	}
}

/*Ф-ия пишет данные в канал раз в секунду*/
func writeToChannel(dataCh chan<- string) {
	/*Запускаем бесконечный цикл*/
	for {
		/*Останавливаем выполнения на одну секунду*/
		time.Sleep(time.Second)
		/*Фиксируем текущее время и отправляем в канал*/
		dataCh <- fmt.Sprintf("%v", time.Now())
	}
}

/*Случшаем прерывание программы*/
func listenInterrupt(dataCh chan string) {
	/*Создаем канал типа Signal*/
	interruptCh := make(chan os.Signal, 1)
	/*signal.Notify ретранслирует сигналы, переданные вторым аргументов,
	в канал, переданный первым аргументом. По сути здесь мы и создали слушатель для SIGINT*/
	signal.Notify(interruptCh, os.Interrupt)

	/*Запускам go routine, который будет ждать данный от канала с прерыванием.
	Когда пользователь нажмет ctrl-C, signal.Notify передаст этот сигнал в interruptCh,
	а эта ф-ия получит данные из канала, и закроет канал с данными*/
	go func() {
		<-interruptCh
		fmt.Println("\nReceived interrupt signal.")
		/*Закрываем канал с данными, тем самым завершая наши воркеры*/
		close(dataCh)
	}()
}
