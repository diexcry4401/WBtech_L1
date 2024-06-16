package main

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterWithMutex struct {
	mutex       sync.Mutex // Мьютекс для защиты критической секции
	countMutex  int        // Счетчик, увеличиваемый с блокировкой мьютекса
	countAtomic int32      // Счетчик, увеличиваемый атомарными операциями
}

func main() {
	var counterMutex CounterWithMutex // Создание экземпляра CounterWithMutex

	var wg sync.WaitGroup // Счетчик для ожидания завершения горутин

	counterMutex.start(&wg) // Запуск метода start с передачей WaitGroup
}

func (c *CounterWithMutex) incrementMutex() {
	c.mutex.Lock()   // Захват мьютекса для исключения доступа других горутин
	c.countMutex++   // Увеличение счетчика под защитой мьютекса
	c.mutex.Unlock() // Освобождение мьютекса после завершения критической секции
}

func (c *CounterWithMutex) incrementAtomic() {
	atomic.AddInt32(&c.countAtomic, 1) // Атомарное увеличение счетчика
}

func (c *CounterWithMutex) start(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ { // Запуск 10 горутин
		wg.Add(1) // Увеличение счетчика активных горутин в WaitGroup

		go func() {
			defer wg.Done() // Уменьшение счетчика активных горутин при завершении

			c.incrementMutex()  // Увеличение счетчика с использованием мьютекса
			c.incrementAtomic() // Увеличение счетчика с использованием атомарной операции
		}()
	}

	wg.Wait() // Ожидание завершения всех запущенных горутин
	fmt.Println("Mutex counter", c.countMutex)
	fmt.Println("Atomic counter", c.countAtomic)
}
