package main

/*
Реализовать собственную функцию sleep.
*/

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Использование различных способов реализации функции sleep
	sleepWithContext(2 * time.Second) // Использование контекста
	sleepWithTimer(2 * time.Second)   // Использование таймера
	sleepWithDelay(2 * time.Second)   // Использование функции After
}

// sleepWithContext использует контекст с таймаутом для реализации функции sleep
func sleepWithContext(duration time.Duration) {
	// Создаем контекст с таймаутом на указанное время
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() // гарантируем что контекст будет отменен

	<-ctx.Done() // по истечению времени приходит сигнал об окончании контекста
	fmt.Printf("Exit after context timeout %s seconds\n", duration)
}

// sleepWithTimer использует таймер для реализации функции sleep
func sleepWithTimer(duration time.Duration) {
	// Создаем новый таймер, который будет отправлять текущее время в свой канал `C` через заданное время `duration`.
	timer := time.NewTimer(duration)

	<-timer.C // Блокируем выполнение функции, ожидая значения из канала таймера `C`.

	// Когда значение из канала `C` получено (то есть, время `duration` прошло), выводим сообщение.
	fmt.Printf("Exit after timer timeout %s seconds\n", duration)
}

// sleepWithDelay использует time.After для реализации функции sleep
func sleepWithDelay(duration time.Duration) {
	/*
			time.After возвращает канал (chan time.Time), который получит текущее время
		    после того, как пройдет указанная продолжительность (duration).
	*/
	<-time.After(duration)

	// Эта строка выполняется после того, как пройдет указанная продолжительность.
	fmt.Printf("Exit after delay timeout %s seconds\n", duration)
}
