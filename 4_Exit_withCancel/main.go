package main

/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые
читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var workers int

	stopSignal, cancel := context.WithCancel(context.Background()) // Создаем контекст, который позволяет завершить программу
	defer cancel()                                                 // Отменяем созданный контекст при выходе из функции
	shutdown(cancel)                                               // Запускаем функцию для завершения программы при получении сигнала Ctrl+C
	fmt.Println("Count of workers:")

	if _, err := fmt.Fscan(os.Stdin, &workers); err != nil || workers <= 0 { // Считываем количество работников из stdin
		fmt.Println("Invalid number of workers")
		return
	}

	start(stopSignal, workers) // Запускаем функцию для запуска работы с указанным количеством работников
}

func start(stopSignal context.Context, workers int) {
	var wg sync.WaitGroup

	data := make(chan int)            // Создаем канал, в который будем записывать данные
	go generateData(stopSignal, data) // Запускаем горутину для генерации данных, передаем канал и контекст для завершения работы горутины

	for i := 0; i < workers; i++ {
		wg.Add(1) // Увеличиваем счетчик для каждой горутины

		go worker(stopSignal, i, data, &wg) // Запускаем воркеров в отдельных горутинах
	}

	defer close(data) // Закрываем канал после завершения работы цикла
	wg.Wait()         // Ожидаем завершения работы всех горутин
}

// Функция воркера, которая читает данные из канала и выводит их в stdout
func worker(stopSignal context.Context, id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик при завершении работы горутины

	for {
		select {
		case <-stopSignal.Done(): // Если пришел сигнал завершения работы (контекст завершения программы), выводим сообщение и завершаем горутину
			fmt.Printf("Worker %d stopping\n", id)
			return
		case number, ok := <-data: // Получаем данные из канала
			if !ok { // Проверяем, не закрылся ли канал
				fmt.Println("Channel is closed")
				return
			}

			fmt.Printf("Worker %d received data: %d\n", id, number) // Выводим полученные данные
		}
	}
}

// Горутина для генерации данных в канал
func generateData(ctx context.Context, data chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // Если пришел сигнал завершения работы (контекст завершения программы), завершаем горутину
			fmt.Println("Stop generating")
			return
		default:
			time.Sleep(time.Second)
			data <- i // Записываем сгенерированные данные в канал для чтения работниками
		}
	}
}

// Функция для завершения программы по получению сигнала Ctrl+C
func shutdown(cancel context.CancelFunc) {
	go func() {
		signalChan := make(chan os.Signal, 1)   // Создаем буферизированный канал для получения сигналов от системы
		signal.Notify(signalChan, os.Interrupt) // Подписываем канал на получение сигнала прерывания (Ctrl+C)
		defer signal.Stop(signalChan)

		<-signalChan // Ожидаем получения сигнала от канала
		cancel()     // Отменяем созданный контекст, что приведет к завершению программы
	}()
}
