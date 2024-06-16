package main

/*
Удалить i-ый элемент из слайса.
*/
import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(numbers) // 2 4 6 8 10
	numbers = deleteWithAppend(numbers, 2)
	fmt.Println(numbers) // 2 4 8 10
	numbers = deleteWithCopy(numbers, 3)
	fmt.Println(numbers) // 2 4 8
}

// deleteWithAppend удаляет элемент из слайса numbers по индексу i с использованием append
func deleteWithAppend(numbers []int, i int) []int {
	return append(numbers[:i], numbers[i+1:]...) // Удаляем элемент с индексом i путем объединения слайсов до и после элемента
}

// deleteWithCopy удаляет элемент из слайса numbers по индексу i с использованием copy и сокращения длины слайса
func deleteWithCopy(numbers []int, i int) []int {
	copy(numbers[i:], numbers[i+1:])   // Копируем элементы от i+1 на место i, сдвигая слайс влево
	numbers = numbers[:len(numbers)-1] // Уменьшаем длину слайса, чтобы удалить дубликат

	return numbers
}
