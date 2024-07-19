package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Crea un programa que comprueba si los paréntesis, llaves y corchetes
de una expresión están equilibrados.
- Equilibrado significa que estos delimitadores se abren y cierran
  en orden y de forma correcta.
- Paréntesis, llaves y corchetes son igual de prioritarios.
  No hay uno más importante que otro.
- Expresión balanceada: { [ a *( c + d ) ] - 5 }
- Expresión no balanceada: { a * ( c + d ) ] - 5 }
*/

func expresionEquilibrada() {

	var expresion string

	fmt.Println("Ingrese la expresión a verificar")

	// Uso de bufio.NewReader para leer la entrada completa del usuario
	reader := bufio.NewReader(os.Stdin)
	expresion, _ = reader.ReadString('\n')

	// Itera sobre las expresiones y verifica si están equilibradas
	if esEquilibrado(expresion) {
		fmt.Printf("La expresión: %sEstá equilibrada \n", expresion)
	} else {
		fmt.Printf("La expresión: %sNo está equilibrada \n", expresion)
	}

}

// Pila define una estructura de pila básica

type Pila struct {
	caracteres []rune
}

// Push agrega un elemento a la pila

func (p *Pila) Push(caracter rune) {
	p.caracteres = append(p.caracteres, caracter)
}

// Pop elimina y devuelve el último elemento de la pila
// Si la pila está vacía, devuelve un valor por defecto y false

func (p *Pila) Pop() (rune, bool) {
	if p.pilaVacia() {

		return ' ', false // Retorna false si la pila está vacía
	}
	ultimo := p.caracteres[len(p.caracteres)-1]       // Obtiene el último elemento
	p.caracteres = p.caracteres[:len(p.caracteres)-1] // Elimina el último elemento del slice

	return ultimo, true
}

// Verifica si la pila está vacía

func (p *Pila) pilaVacia() bool {
	return len(p.caracteres) == 0
}

// esEquilibrado verifica si una expresión tiene los delimitadores equilibrados

func esEquilibrado(expresion string) bool {

	pila := &Pila{} // Inicializa una nueva pila

	for _, caracter := range expresion {

		// Si el carácter es un delimitador de apertura, lo añade a la pila

		if caracter == '(' || caracter == '{' || caracter == '[' {
			pila.Push(caracter)
		} else if caracter == ')' || caracter == '}' || caracter == ']' {

			// Si es un delimitador de cierre, verifica si la pila está vacía

			if pila.pilaVacia() {
				return false // Hay un delimitador de cierre sin su correspondiente apertura
			}
			top, _ := pila.Pop() // Saca el último elemento de la pila
			if !sonCorrespondientes(top, caracter) {
				return false // Los delimitadores no coinciden
			}
		}
	}
	// Verifica si la pila está vacía al final
	return pila.pilaVacia()
}

// sonCorrespondientes verifica si los delimitadores de apertura y cierre coinciden

func sonCorrespondientes(abrir, cerrar rune) bool {
	return (abrir == '(' && cerrar == ')') ||
		(abrir == '{' && cerrar == '}') ||
		(abrir == '[' && cerrar == ']')
}
