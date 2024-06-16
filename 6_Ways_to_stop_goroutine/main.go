package main

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	exitWithContext() // Выход из горутины с использованием контекста
	fmt.Println("--------")
	exitWithChannel() // Выход из горутины с использованием закрытия канала
	fmt.Println("--------")
	exitWithTimer() // Выход из горутины с использованием таймера
	fmt.Println("--------")
	exitWithAtomicVar() // Выход из горутины с использованием атомарной переменной
}

func exitWithContext() {
	ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст, позволяющий завершить выполнение программы
	defer cancel()                                          // Отменяем контекст при выходе из функции

	var wg sync.WaitGroup

	wg.Add(1) // Увеличиваем счетчик для горутины

	go func() { // Запускаем горутину
		defer wg.Done() // Уменьшаем счетчик при завершении горутины

		for {
			select {
			case <-ctx.Done(): // Получаем сигнал о завершении работы контекста
				fmt.Println("Exit with context")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("Goroutine is working")
			}
		}
	}()

	time.Sleep(time.Second * 2) // Даем горутине поработать 2 секунды перед завершением
	cancel()                    // Отменяем контекст, что приведет к завершению работы горутины
	wg.Wait()                   // Ожидаем завершения работы всех горутин
}

func exitWithChannel() {
	semaphore := make(chan struct{}) // Создаем канал для сигнализации завершения работы горутины

	var wg sync.WaitGroup

	wg.Add(1) // Увеличиваем счетчик для горутины

	go func() { // Запускаем горутину
		defer wg.Done() // Уменьшаем счетчик при завершении горутины

		for {
			select {
			case <-semaphore: // Получаем сигнал о закрытии канала
				fmt.Println("Exit with closed channel")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("Goroutine is working")
			}
		}
	}()

	time.Sleep(time.Second * 2) // Даем горутине поработать 2 секунды перед завершением
	close(semaphore)            // Закрываем канал, что приведет к завершению работы горутины
	wg.Wait()                   // Ожидаем завершения работы всех горутин
}

func exitWithTimer() {
	var wg sync.WaitGroup

	timer := time.NewTimer(time.Second * 2) // Создаем таймер на 2 секунды

	wg.Add(1) // Увеличиваем счетчик для горутины

	go func(timer <-chan time.Time) { // Запускаем горутину, передаем канал таймера
		defer wg.Done() // Уменьшаем счетчик при завершении горутины

		for {
			select {
			case <-timer: // Получаем сигнал об истечении времени таймера
				fmt.Println("Exit with timer")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("Goroutine is working")
			}
		}
	}(timer.C)

	wg.Wait() // Ожидаем завершения работы всех горутин
}

func exitWithAtomicVar() {
	var run int32 = 1 // Атомарная переменная, определяющая состояние работы горутины

	var wg sync.WaitGroup

	wg.Add(1) // Увеличиваем счетчик для горутины

	go func() { // Запускаем горутину
		defer wg.Done() // Уменьшаем счетчик при завершении горутины

		for atomic.LoadInt32(&run) == 1 { // Проверяем состояние атомарной переменной
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Goroutine is working")
		}
		fmt.Println("Exit with atomic")
	}()

	time.Sleep(time.Second * 2) // Даем горутине поработать 2 секунды перед завершением
	atomic.StoreInt32(&run, 0)  // Устанавливаем атомарную переменную в 0, что приведет к завершению работы горутины
	wg.Wait()                   // Ожидаем завершения работы всех горутин
}
