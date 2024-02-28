package main

import (
	"fmt"
)

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Какое действие вы хотите выполнить? (+, -, *, /)")

	var action string
	fmt.Scan(&action)

	var a float64
	var b float64

	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Println("Введите второе число: ")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("Результат сложения: " + fmt.Sprint(a+b))

	case action == "-":
		fmt.Println("Результат убавления: = " + fmt.Sprint(a-b))

	case action == "*":
		fmt.Println("Результат умножения: = " + fmt.Sprint(a*b))

	case action == "/":
		fmt.Println("Результат деления: = " + fmt.Sprint(a/b))
	}
}
