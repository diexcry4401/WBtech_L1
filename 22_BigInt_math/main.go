package main

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/
import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1234567890123456789)
	b := big.NewInt(4401440144014401)

	fmt.Println(new(big.Int).Add(a, b)) // Cложение
	fmt.Println(new(big.Int).Div(a, b)) // Вычитание
	fmt.Println(new(big.Int).Mul(a, b)) // Умножение
	fmt.Println(new(big.Int).Sub(a, b)) // Вычитание
}
