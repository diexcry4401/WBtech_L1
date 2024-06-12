package main

/*
  Дана структура Human (с произвольным набором полей и методов).
  Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

import "fmt"

// Human - родительская структура с полями Name и Age
type Human struct {
	Name string
	Age  int
}

// printName - метод структуры Human, который выводит имя
func (h *Human) printName() {
	fmt.Println("Name:", h.Name)
}

// printAge - метод структуры Human, который выводит возраст
func (h *Human) printAge() {
	fmt.Println("Age:", h.Age)
}

// Action - структура, в которую встроена структура Human
// Это позволяет структуре Action наследовать поля и методы структуры Human
type Action struct {
	Human
}

// addYear - метод структуры Action, который увеличивает возраст на 1
// и выводит новое значение возраста
func (a *Action) addYear() {
	a.Age += 1
	fmt.Println("Now age is:", a.Age) // вывод увеличенной переменной Age из структуры Action, унаследованной от Human
}

func main() {
	name := "Max"
	age := 23

	// Создаем экземпляр структуры Action и инициализируем встроенную структуру Human
	Teacher := Action{
		Human{
			Name: name,
			Age:  age,
		},
	}

	// Вызываем методы структуры Human через экземпляр структуры Action
	Teacher.printName() // Выводит имя
	Teacher.printAge()  // Выводит возраст
	Teacher.addYear()   // Увеличивает возраст на 1 и выводит новое значение возраста
}
