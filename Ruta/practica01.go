//* Operadores y estructuras de control

package main

import "fmt"

// DIFICULTAD EXTRA
// Crea un programa que imprima por consola todos los números comprendidos entre 10 y 55 (incluidos), pares, y que no son ni el 16 ni múltiplos de 3. Seguro que al revisar detenidamente las posibilidades has descubierto algo nuevo.

func practica01() /*main*/ {

	// https://go.dev/doc/

	for i := 10; i <= 55; i++ {

		if i%2 == 0 && i != 16 && i%3 != 0 {
			fmt.Println(i)
		}
	}

}
