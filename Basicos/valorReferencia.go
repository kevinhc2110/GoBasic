package main

import "fmt"

// Asignación por valor:

func asignacionValor() {
    // Asignación por valor para tipos básicos
    x := 10
    y := x
    fmt.Println("x:", x, "y:", y) // Imprime: x: 10 y: 10

    // Modificar 'y' no afecta a 'x'
    y = 20
    fmt.Println("x:", x, "y:", y) // Imprime: x: 10 y: 20
}

// Asignación por referencia:

type Persona3 struct {
	Nombre string
	Edad   int
}

func asignacionReferencia() {

	// Asignación por referencia para tipos de referencia (como structs)

	persona4 := Persona3{"Juan", 30}
	persona5 := &persona4
	fmt.Println("persona4:", persona4, "persona5:", *persona5) // Imprime: persona1: {Juan 30} persona2: {Juan 30}

	// Modificar 'persona2' afecta a 'persona1'

	persona5.Nombre = "Pedro"
	fmt.Println("persona4:", persona4, "persona5:", *persona5) // Imprime: persona1: {Pedro 30} persona2: {Pedro 30}
}

// Funciones con paso por valor

func duplicarValor(numero int) {
	numero *= 2
}

func fPasoValor() {
	x := 10
	duplicarValor(x)
	fmt.Println("x después de duplicarValor:", x) // Imprime: x después de duplicarValor: 10
}

// Funciones con paso por referencia

func duplicarReferencia(numero *int) {
	*numero *= 2
}

func fPasoReferencia() {
	x := 10
	duplicarReferencia(&x)
	fmt.Println("x después de duplicarReferencia:", x) // Imprime: x después de duplicarReferencia: 20
}

// La asignación por valor crea una copia independiente del valor original, mientras que la asignación por referencia crea un alias que apunta al mismo valor en la memoria.
// Las variables básicas como enteros y cadenas se asignan por valor, mientras que las estructuras de datos complejas como los slices y los maps se asignan por referencia.
// Las funciones en Go pasan argumentos por valor de forma predeterminada, pero pueden recibir punteros para simular un paso por referencia, permitiendo así modificar el valor original dentro de la función.