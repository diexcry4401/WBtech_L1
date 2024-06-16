package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/
import "fmt"

func main() {
	original := "горутина, голанг"      // Оригинальная строка, которую нужно перевернуть
	reversed := reverseString(original) // Вызываем функцию для переворота строки
	fmt.Println(reversed)               // Выводим перевернутую строку
}

func reverseString(str string) string {
	reversed := []rune(str) // Преобразуем строку в слайс рун (Unicode символов)
	strLen := len(reversed) // Получаем длину строки в рунах

	// Обходим строку с обеих сторон к центру
	for i, j := 0, strLen-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i] // Меняем местами символы в слайсе рун
	}

	return string(reversed) // Преобразуем слайс рун обратно в строку и возвращаем результат
}
