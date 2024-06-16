package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для передачи значений
	ch := make(chan int)

	// Вводим количество секунд работы программы
	var duration int
	fmt.Print("Enter duration in seconds: ")
	fmt.Scan(&duration)

	// Горутина для отправки значений в канал
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(duration) * time.Second):
				close(ch)
				return
			default:
				ch <- i
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Горутина для чтения значений из канала
	go func() {
		for val := range ch {
			fmt.Println("Received:", val)
		}
		fmt.Println("Channel closed")
	}()

	// Ожидаем завершения программы
	time.Sleep(time.Duration(duration+1) * time.Second)
	fmt.Println("Program finished")
}
