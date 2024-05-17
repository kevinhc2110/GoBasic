package main

import "fmt"

func operadores /*main*/() {
    var a = 10
    var b = 5

		// Operadores Aritméticos

    fmt.Println("a + b =", a + b)  // Suma
    fmt.Println("a - b =", a - b)  // Resta
    fmt.Println("a * b =", a * b)  // Multiplicación
    fmt.Println("a / b =", a / b)  // División
    fmt.Println("a % b =", a % b)  // Módulo (resto de la división)

		//Operadores de Comparación

    fmt.Println("a == b =", a == b)  // Igual a
    fmt.Println("a != b =", a != b)  // Diferente a
    fmt.Println("a > b =", a > b)    // Mayor que
    fmt.Println("a < b =", a < b)    // Menor que
    fmt.Println("a >= b =", a >= b)  // Mayor o igual que
    fmt.Println("a <= b =", a <= b)  // Menor o igual que

		//Operadores lógicos

		var c = true
    var d = false

		fmt.Println("c && d =", c && d)  // AND lógico
    fmt.Println("c || d =", c || d)  // OR lógico
    fmt.Println("!c =", !c)          // NOT lógico

		//Operadores Bit a Bit

		fmt.Println("a & b =", a & b)    // AND bit a bit
    fmt.Println("a | b =", a | b)    // OR bit a bit
    fmt.Println("a ^ b =", a ^ b)    // XOR bit a bit
    fmt.Println("a &^ b =", a &^ b)  // AND NOT bit a bit
    fmt.Println("a << 1 =", a << 1)  // Desplazamiento a la izquierda
    fmt.Println("a >> 1 =", a >> 1)  // Desplazamiento a la derecha

		//Operadores de Asignación

		a += 5  // a = a + 5
    fmt.Println("a += 5 ->", a)

    a -= 3  // a = a - 3
    fmt.Println("a -= 3 ->", a)

    a *= 2  // a = a * 2
    fmt.Println("a *= 2 ->", a)

    a /= 4  // a = a / 4
    fmt.Println("a /= 4 ->", a)

    a %= 3  // a = a % 3
    fmt.Println("a %= 3 ->", a)

    a <<= 1 // a = a << 1
    fmt.Println("a <<= 1 ->", a)

    a >>= 1 // a = a >> 1
    fmt.Println("a >>= 1 ->", a)

    a &= 2  // a = a & 2
    fmt.Println("a &= 2 ->", a)

    a |= 1  // a = a | 1
    fmt.Println("a |= 1 ->", a)

    a ^= 3  // a = a ^ 3
    fmt.Println("a ^= 3 ->", a)

    a &^= 1 // a = a &^ 1
    fmt.Println("a &^= 1 ->", a)


		// Operadores especiales
		//Operador de Dirección (&): Obtiene la dirección de memoria de una variable.
		//Operador de Desreferencia (*): Accede al valor apuntado por un puntero.

		var p *int

    p = &a  // p obtiene la dirección de a
    fmt.Println("Dirección de a:", p)
    fmt.Println("Valor de a a través de p:", *p)  // Desreferencia de p


		// Desreferencia: Se refiere al proceso de acceder al valor almacenado en la dirección de memoria a la que apunta un puntero
}