package main

import (
	"fmt"
	"math/big"
)


func primon() {

	for numero := 1; numero <= 100; numero++ {
		if esPrimo(numero) {
			fmt.Println(numero)
		}
	}

	for numero := 1; numero <= 100; numero++ {
		if esPrimo2(numero) {
			fmt.Println(numero)
		}
	}

	primos := cribaEratostenes(100)
    for _, primo := range primos {
        fmt.Println(primo)
    }
}


// Opción 1

func esPrimo(numero int) bool {

	if numero <= 1 {
		return false
	}
	if numero <= 3 {
		return true
	}
	if numero%2 == 0 || numero%3 == 0 {
		return false
	}

	// Inicializa i con el valor 5. Aquí i es el primer número impar mayor que 3 que necesitamos verificar.

	i := 5

	// Este bucle continúa mientras i * i sea menor o igual que numero.

	for i*i <= numero {

		// Esta verificación funciona porque todos los números primos mayores que 3 pueden ser expresados como 6𝑘±1 para algún entero k. Esto excluye los múltiplos de 2 y 3:

		if numero%i == 0 || numero%(i+2) == 0 {
			return false
		}

		// Esto significa que después de verificar i y i + 2 (por ejemplo, 5 y 7), se pasa a los próximos candidatos, que son i + 6 y (i + 6) + 2 (por ejemplo, 11 y 13).

		i += 6
	}
	return true
}

// Opción 2

func esPrimo2(numero int) bool {

	// Convertir el número a tipo big.Int

	bigNumero := big.NewInt(int64(numero))

	// Verificar si el número es probablemente primo

	return bigNumero.ProbablyPrime(0) // 0 iterations implies deterministic check for small numbers
}

// Opción  3

func cribaEratostenes(limite int) []int {

	// Se crea un slice de booleanos llamado primos de longitud limite+1.
	// los elementos del slice se inicializan como true para marcarlos inicialmente como primos.
	// Los índices 0 y 1 se establecen como false ya que 0 y 1 no son números primos.

	primos := make([]bool, limite+1)

	for i := range primos {
			primos[i] = true
	}
	primos[0] = false
	primos[1] = false

	// Para cada número i desde 2 hasta la raíz cuadrada de limite
	// Entonces, para cada múltiplo de i desde i*i hasta limite, se marca ese múltiplo como false en el slice primos, ya que los múltiplos de i no pueden ser primos.

	for i := 2; i*i <= limite; i++ {
			if primos[i] {
					for j := i * i; j <= limite; j += i {
							primos[j] = false
					}
			}
	}

	// Después de que se han marcado todos los múltiplos de los números primos como false, los números que quedan marcados como true en el slice primos son los números primos dentro del rango.
	// Se recorre el slice primos nuevamente y se agregan los índices que están marcados como true al slice resultado.

	var resultado []int
	for i := 2; i <= limite; i++ {
			if primos[i] {
					resultado = append(resultado, i)
			}
	}
	return resultado
}