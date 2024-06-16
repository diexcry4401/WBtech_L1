package main

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/
import "fmt"

func main() {
	variable1 := 1
	variable2 := "string"
	variable3 := true
	variable4 := make(chan int)
	variable5 := make(chan string)
	variable6 := make(chan bool)
	variable7 := 13.37

	// Вызов функции typeIdentifier с различными переменными разных типов данных
	typeIdentifier(variable1, variable2, variable3, variable4, variable5, variable6, variable7)
}

func typeIdentifier(variables ...interface{}) {
	for _, variable := range variables { // Итерация по переданным переменным в функцию
		switch expr := variable.(type) { // Определение типа переменной
		case int:
			fmt.Println("variable is int", expr) // Вывод информации о типе переменной и ее значении (если int)
		case string:
			fmt.Println("variable is string", expr) // Вывод информации о типе переменной и ее значении (если string)
		case bool:
			fmt.Println("variable is bool", expr) // Вывод информации о типе переменной и ее значении (если bool)
		case chan int:
			fmt.Println("variable is chan int", expr) // Вывод информации о типе переменной и ее значении (если канал int)
		case chan string:
			fmt.Println("variable is chan string", expr) // Вывод информации о типе переменной и ее значении (если канал string)
		case chan bool:
			fmt.Println("variable is chan bool", expr) // Вывод информации о типе переменной и ее значении (если канал bool)
		default:
			fmt.Println("invalid type", expr) // Вывод сообщения об ошибочном типе переменной (если не подошел ни один case)
		}
	}
}
