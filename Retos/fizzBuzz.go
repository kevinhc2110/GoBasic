package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Escribe un programa que muestre por consola (con un print) los
números de 1 a 100 (ambos incluidos y con un salto de línea entre
cada impresión), sustituyendo los siguientes:
- Múltiplos de 3 por la palabra "fizz".
- Múltiplos de 5 por la palabra "buzz".
- Múltiplos de 3 y de 5 a la vez por la palabra "fizzbuzz".
*/

func fizzbuzz() {

	// Opción 1

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Opción 2 String.Builder

	var result strings.Builder

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			result.WriteString("Fizz")
		}
		if i%5 == 0 {
			result.WriteString("Buzz")
		}
		if result.Len() == 0 {
			result.WriteString(strconv.Itoa(i))
		}
		if result.Len() > 0 {
			println(result.String())
			result.Reset()
		}
	}
}
