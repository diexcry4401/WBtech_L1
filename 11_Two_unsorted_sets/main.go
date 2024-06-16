package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/
import "fmt"

func main() {
	var commonData []string // Слайс для хранения элементов пересечения множеств

	// Создание двух множеств
	set1 := initSet("bmw", "ferrari", "cadillac", "porsche", "mustang")
	set2 := initSet("alphine", "bmw", "porsche", "mini")

	for elem := range set1 { // Итерация по элементам первого множества
		if _, contains := set2[elem]; contains { // Проверка наличия текущего элемента во втором множестве
			commonData = append(commonData, elem) // Добавление элемента в слайс пересечений
		}
	}

	fmt.Println(commonData) // Вывод слайса с пересечениями
}

func initSet(keys ...string) map[string]struct{} {
	set := make(map[string]struct{}) // Инициализация пустого множества

	for _, key := range keys { // Итерация по переданным ключам
		set[key] = struct{}{} // Добавление ключа в множество со значением struct{}
	}

	return set // Возврат готового множества
}
