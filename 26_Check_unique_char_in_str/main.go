package main

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	string1 := "abcd"
	string2 := "abCdefAaf"
	string3 := "aabcd"

	fmt.Println(ValidationUniqSymbolsWithSort(string1))
	fmt.Println(ValidationUniqSymbolsWithSort(string2))
	fmt.Println(ValidationUniqSymbolsWithSort(string3))
	fmt.Println("--------")
	fmt.Println(ValidationUniqSymbolsWithSet(string1))
	fmt.Println(ValidationUniqSymbolsWithSet(string2))
	fmt.Println(ValidationUniqSymbolsWithSet(string3))
	fmt.Println("--------")
}

func ValidationUniqSymbolsWithSort(str string) bool {
	str = strings.ToLower(str)         // Приводим все символы строки к нижнему регистру
	sliceStr := strings.Split(str, "") // Разбиваем строку на слайс
	sort.Strings(sliceStr)             // Сортируем слайс, гарантируя, что если будут одинаковые символы, то они будут стоять рядом
	for i := 1; i < len(sliceStr); i++ {
		if sliceStr[i] == sliceStr[i-1] { // Со индексу проверяем, если символ равен предидущему - выходим
			return false
		}
	}

	return true
}

func ValidationUniqSymbolsWithSet(str string) bool {
	str = strings.ToLower(str)        // Приводим все символы строки к нижнему регистру
	setStr := make(map[rune]struct{}) // Создаем map для хранения уникальных символов (используем как set)
	for _, char := range str {
		if _, ok := setStr[char]; ok { // Проверяем, существует ли уже символ в set
			return false // Если символ уже есть в set, возвращаем false (символ не уникален)
		} else {
			setStr[char] = struct{}{} // Если символа нет в set, добавляем его
		}
	}
	return true // Если все символы уникальны (добавились в сет), возвращаем true
}
