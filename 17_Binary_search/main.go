package main

/*
Реализовать бинарный поиск встроенными методами языка.
*/
import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10, 11}

	// Вызываем функцию binarySearch для поиска чисел в отсортированном массиве
	number := binarySearch(numbers, 4) // ищем число 4
	fmt.Println(number)                // выводим индекс найденного числа (1)

	number = binarySearch(numbers, 10) // ищем число 10
	fmt.Println(number)                // выводим индекс найденного числа (4)

	number = binarySearch(numbers, 9) // ищем число 9 (которого нет в массиве)
	fmt.Println(number)               // выводим -1, так как число 9 не найдено
}

func binarySearch(numbers []int, searchedNumber int) int {
	left, right := 0, len(numbers)-1 // Определяем границы для поиска в массиве

	for left <= right { // Пока левая граница не превысит правую
		middle := (left + right) / 2 // Находим индекс середины массива

		switch {
		case searchedNumber < numbers[middle]:
			right = middle - 1 // Искомое число меньше числа в середине, сдвигаем правую границу
		case searchedNumber > numbers[middle]:
			left = middle + 1 // Искомое число больше числа в середине, сдвигаем левую границу
		default:
			return middle // Возвращаем индекс найденного числа
		}
	}

	return -1 // Возвращаем -1, если искомый элемент не найден
}
