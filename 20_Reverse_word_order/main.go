package main

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
import (
	"fmt"
	"strings"
)

func main() {
	original := "snow dog sun"                 // Исходная строка, содержащая слова для обработки
	reversed := reverseWordsInString(original) // Вызываем функцию для обратного порядка слов
	fmt.Println(original)                      // Выводим результат оригинал
	fmt.Println(reversed)                      // Выводим результат обработки
}

func reverseWordsInString(str string) string {
	words := strings.Fields(str) // Разбиваем входную строку на слайс слов
	strLen := len(words)         // Получаем количество слов в строке

	// Обходим слайс слов с обеих сторон к центру
	for i, j := 0, strLen-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем местами слова в слайсе
	}

	return strings.Join(words, " ") // Склеиваем слова в строку с пробелами между ними и возвращаем результат
}
