package main

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

func main() {

	resultado := Sumar(3, 4)
	fmt.Println("La suma de 3 y 4 es:", resultado)
}