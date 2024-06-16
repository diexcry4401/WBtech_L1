package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Непредсказуемое поведение: Если someFunc вызывается параллельно из разных потоков, это может привести к конфликтам доступа к данным (гонкам данных),
что приведет к непредсказуемому поведению или потере данных в переменной justString.

Утечка памяти: Использование justString[:len] создает срез, который ссылается на весь изначальный массив данных.
Это означает, что массив не может быть освобожден из памяти, пока существует хотя бы одна ссылка на него.
*/
import "fmt"

func createHugeString(stringSize int) string {
	byteArray := make([]byte, stringSize) // Создание слайса байтов заданного размера
	for i := 0; i < stringSize; i++ {
		byteArray[i] = 'X' // Заполнение слайса байтов символами 'X'
	}
	return string(byteArray) // Преобразование байтового массива в строку
}

func someFunc(justString string, length int) string {
	return string([]rune(justString)[:length]) // Создание новой подстроки из исходной строки
}

func main() {
	justString := createHugeString(1 << 10) // Создание большой строки
	newString := someFunc(justString, 5)    // Получение подстроки длиной 5 символов
	fmt.Println(newString)                  // Вывод результата
}
