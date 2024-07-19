package main

import (
	"fmt"
	"strconv"
)

/*
Escribe una función que calcule si un número dado es un número de Armstrong
(o también llamado narcisista).
Si no conoces qué es un número de Armstrong, debes buscar información
al respecto.
*/

func esArmstrong() {

	num := 54748

	// Opción 1

	if numeroArmstrong1(num) {
		fmt.Printf("El número %d es un número de Armstrong.\n", num)
	} else {
		fmt.Printf("El número %d no es un número de Armstrong.\n", num)
	}

	// Opción 2

	if numeroArmstrong2(num) {
		fmt.Printf("El número %d es un número de Armstrong.\n", num)
	} else {
		fmt.Printf("El número %d no es un número de Armstrong.\n", num)
	}
}

// Opción 1

func numeroArmstrong1(num int) bool {

	// Convertir entero a string

	numStr := strconv.Itoa(num)
	potencia := len(numStr)
	var verificado int

	for _, numero := range numStr {

		digitoStr := string(numero)

		digito, err := strconv.Atoi(digitoStr)

		if err != nil {
			fmt.Println("Error", err)
			return false
		}

		// Opción 1

		// verificado += int(math.Pow(float64(digito), float64(potencia)))

		// Opción 2

		verificado += potenciaEntera(digito, potencia)
	}

	return num == verificado

}

// Opción 2

func potenciaEntera(base, exponente int) int {
	resultado := 1
	for i := 0; i < exponente; i++ {

		// Multiplicamos la base por si mismo
		resultado *= base
	}
	return resultado
}

// Opción 3

func numeroArmstrong2(num int) bool {
	// Calcula el número de dígitos en el número
	numDigitos := contarDigitos(num)
	// Copia del número original para calcular la suma
	restante := num
	var suma int

	// Calcula la suma de los dígitos elevados a la potencia
	for restante > 0 {
		digito := restante % 10
		suma += potencia(digito, numDigitos)
		restante /= 10
	}

	// Verifica si el número es un número de Armstrong
	return num == suma
}

func contarDigitos(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func potencia(base, exponente int) int {
	resultado := 1
	for i := 0; i < exponente; i++ {
		resultado *= base
	}
	return resultado
}
