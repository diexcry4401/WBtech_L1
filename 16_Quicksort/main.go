package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
import "fmt"

func main() {
	numbers := []int{8, 7, 9, 14, 2, 23, 16, 4, 8}
	numbers = quickSort(numbers) // Сортируем массив чисел с помощью quickSort
	fmt.Println(numbers)         // Выводим отсортированный массив
}

func quickSort(numbers []int) []int {
	if len(numbers) < 2 { // Базовый случай: если в массиве 1 элемент или пусто, он считается отсортированным
		return numbers
	}

	var less, greater []int // Слайсы для значений меньших и больших, чем первый элемент (pivot)

	elem := numbers[0]                   // Выбираем первый элемент массива как опорный элемент (pivot)
	for _, number := range numbers[1:] { // Проходим по оставшимся элементам массива
		if number <= elem {
			less = append(less, number) // Если элемент меньше или равен опорному, добавляем его в слайс less
		} else {
			greater = append(greater, number) // Если элемент больше опорного, добавляем его в слайс greater
		}
	}

	// Рекурсивно сортируем слайсы less и greater
	result := append(quickSort(less), elem)        // Сортируем и объединяем less с опорным элементом
	result = append(result, quickSort(greater)...) // Сортируем и объединяем result с отсортированным greater

	return result // Возвращаем отсортированный массив
}
