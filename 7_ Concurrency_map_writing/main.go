package main

/*
Реализовать конкурентную запись данных в map.
*/
import (
	"fmt"
	"sync"
)

func main() {
	sliceNumbers := []int{2, 4, 6, 8, 10}
	writeWithMutex(sliceNumbers) // Запускаем функцию записи с использованием мьютекса
	fmt.Println("--------")
	writeWithSyncMap(sliceNumbers) // Запускаем функцию записи с использованием sync.Map
}

func writeWithMutex(sliceNumbers []int) {
	var wg sync.WaitGroup           // Создаем WaitGroup для ожидания завершения всех горутин
	var mutex sync.Mutex            // Создаем Mutex для защиты доступа к map
	mapNumbers := make(map[int]int) // Создаем обычный map для хранения данных

	for index, number := range sliceNumbers {
		wg.Add(1) // Увеличиваем счетчик горутин в WaitGroup

		go func(index, number int) {
			defer wg.Done() // Уменьшаем счетчик горутин при завершении

			mutex.Lock()               // Захватываем мьютекс перед записью в map
			mapNumbers[index] = number // Записываем значение в map
			mutex.Unlock()             // Освобождаем мьютекс после записи
		}(index, number)
	}

	wg.Wait() // Ожидаем завершения всех горутин

	fmt.Println("Map with Mutex:", mapNumbers) // Выводим map после всех операций
}

func writeWithSyncMap(sliceNumbers []int) {
	var wg sync.WaitGroup   // Создаем WaitGroup для ожидания завершения всех горутин
	var mapNumbers sync.Map // Создаем sync.Map для безопасной записи данных

	for index, number := range sliceNumbers {
		wg.Add(1) // Увеличиваем счетчик горутин в WaitGroup

		go func(index, number int) {
			defer wg.Done() // Уменьшаем счетчик горутин при завершении

			mapNumbers.Store(index, number) // Записываем значение в sync.Map
		}(index, number)
	}

	wg.Wait() // Ожидаем завершения всех горутин

	mapNumbers.Range(func(key, value interface{}) bool { // Итерируемся по всем элементам sync.Map
		fmt.Println(key, value) // Выводим ключ и значение на экран
		return true             // Возвращаем true, чтобы продолжить итерацию
	})
}
