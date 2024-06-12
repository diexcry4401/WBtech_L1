package main

import (
	"fmt"
	"sync"
)

// Функция для вычисления квадрата числа
func square(num int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done() // Уменьшаем счетчик горутин в конце выполнения функции
	result := num * num
	results <- result // Отправляем результат в канал
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}        // Исходный массив чисел
	results := make(chan int, len(numbers)) // Канал для результатов
	var wg sync.WaitGroup                   // WaitGroup для ожидания завершения всех горутин

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)                    // Увеличиваем счетчик горутин
		go square(num, &wg, results) // Запускаем горутину функции вычесления квадрата
	}

	// Ждем завершения всех горутин
	wg.Wait()
	close(results) // Закрываем канал после завершения всех горутин

	// Читаем результаты из канала и выводим их
	for result := range results {
		fmt.Println(result)
	}
}
