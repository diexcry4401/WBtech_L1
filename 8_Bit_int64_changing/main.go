package main

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
import "fmt"

func main() {
	var number int64 = 5 // Исходное число 5 в бинарном виде: 0101

	i := 1                                    // Позиция бита, который будем устанавливать или сбрасывать (1 в бинарном виде: 0001)
	newNumber := setBitsInInt64(number, i, 1) // Устанавливаем бит в позиции i в 1: после побитового ИЛИ получаем 0111
	fmt.Println(newNumber)

	i = 2                                    // Обновляем позицию бита (2 в бинарном виде: 0010)
	newNumber = setBitsInInt64(number, i, 0) // Сбрасываем бит в позиции i: после побитового И получаем 0001
	fmt.Println(newNumber)
}

func setBitsInInt64(number int64, i int, bit int) int64 {
	if bit == 0 {
		return clearBit(number, i) // Если bit равен 0, сбрасываем бит
	}

	return setBit(number, i) // Если bit равен 1, устанавливаем бит
}

func clearBit(number int64, i int) int64 {
	return number &^ (1 << i)
} // Функция clearBit сбрасывает бит в позиции i числа number с помощью побитовой инверсии и логического И.

func setBit(number int64, i int) int64 {
	return number | (1 << i)
} // Функция setBit устанавливает бит в позиции i числа number с помощью побитового ИЛИ.
