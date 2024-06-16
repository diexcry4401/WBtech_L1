package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
import "fmt"

func main() {
	consistency := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{}) // Инициализация пустого множества

	for _, elem := range consistency { // Итерация по переданным элементам в последовательности
		set[elem] = struct{}{} // Добавление элемента в множество с пустым значением struct{}
	}

	fmt.Println(set) // Вывод множества в консоль
}
