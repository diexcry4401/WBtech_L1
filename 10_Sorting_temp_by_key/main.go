package main

/*
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
import (
	"fmt"
)

func main() {
	consistency := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1, 0}

	subsets := make(map[int][]float64) // Создание карты для хранения подмножеств температур

	for _, temperature := range consistency { // Итерация по всем температурам из последовательности
		key := int(temperature/10.0) * 10                // Вычисление ключа для текущей температуры
		subsets[key] = append(subsets[key], temperature) // Добавление текущей температуры в соответствующее подмножество
	}

	fmt.Println(subsets) // Вывод результата - карты с группами температур
}
