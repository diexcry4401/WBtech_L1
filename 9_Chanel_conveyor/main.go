package main

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	chan1 := make(chan int) // Первый канал в который пишутся числа (x) из массива
	chan2 := make(chan int) // Второй канал в который пишутся результаты операции x*2

	go writeInFirstChan(numbers, chan1) // Запускаем горутину для записи слайса в 1ый канал

	go mul2FromFirstChanInSecond(chan1, chan2) // Запускаем горутину для записи чисел *2 во 2ой канал

	for number := range chan2 { // Считываем числа из второго канала
		fmt.Println(number)
	}
}

func writeInFirstChan(numbers []int, first chan int) {
	for _, number := range numbers {
		first <- number // Записываем числа из слайса в канал
	}

	close(first) // Закрываем канал
}

func mul2FromFirstChanInSecond(first chan int, second chan int) {
	for number := range first {
		second <- number * 2 // Умножаем числа на 2 из первого канала и записываем во второй канал
	}

	close(second) // закрываем канал
}
