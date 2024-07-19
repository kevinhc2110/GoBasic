package main

import "fmt"

/*
Escribe una función que calcule y retorne el factorial de un número dado
de forma recursiva.
*/

func factorialRecursivo() {

	fmt.Println(factorial(5))
}

func factorial(num int) int {

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
