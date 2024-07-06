package main

import (
	"fmt"
	"math"
)

/* DIFICULTAD EXTRA:
Desarrolla una calculadora que necesita realizar diversas operaciones matemáticas.
Requisitos:
- Debes diseñar un sistema que permita agregar nuevas operaciones utilizando el OCP.
Instrucciones:
1. Implementa las operaciones de suma, resta, multiplicación y división.
2. Comprueba que el sistema funciona.
3. Agrega una quinta operación para calcular potencias.
4. Comprueba que se cumple el OCP.
*/

// Operation define una interfaz para operaciones matemáticas
type Operation interface {
	Calculate(a, b float64) float64
}

// Add implementa la operación de suma
type Add struct{}

func (o Add) Calculate(a, b float64) float64 {
	return a + b
}

// Subtract implementa la operación de resta
type Subtract struct{}

func (o Subtract) Calculate(a, b float64) float64 {
	return a - b
}

// Multiply implementa la operación de multiplicación
type Multiply struct{}

func (o Multiply) Calculate(a, b float64) float64 {
	return a * b
}

// Divide implementa la operación de división
type Divide struct{}

func (o Divide) Calculate(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	fmt.Println("Error: División por cero")
	return math.NaN()
}

// Calculator procesa operaciones matemáticas
type Calculator struct{}

func (c Calculator) PerformOperation(op Operation, a, b float64) float64 {
	return op.Calculate(a, b)
}

func practica27() {

	calculator := Calculator{}
	add := Add{}
	subtract := Subtract{}
	multiply := Multiply{}
	divide := Divide{}
	// Se implementa sin modificar potencia
	power := Power{}

	fmt.Println("Suma:", calculator.PerformOperation(add, 10, 5))
	fmt.Println("Resta:", calculator.PerformOperation(subtract, 10, 5))
	fmt.Println("Multiplicación:", calculator.PerformOperation(multiply, 10, 5))
	fmt.Println("División:", calculator.PerformOperation(divide, 10, 5))
	fmt.Println("Potencia:", calculator.PerformOperation(power, 2, 3)) // 2^3 = 8
}

// Power implementa la operación de potencia
type Power struct{}

func (o Power) Calculate(a, b float64) float64 {
	return math.Pow(a, b)
}