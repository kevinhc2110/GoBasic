package main

import "fmt"

//Las pilas y las colas son estructuras de datos fundamentales en informática que se utilizan para organizar elementos de manera específica y eficiente en función de cómo se accede y se elimina cada elemento de la estructura.

// Pila (Stack - LIFO)
// El último elemento que se añade a la pila es el primero en ser eliminado.

// Definición de la estructura Stack

type Stack []int

// Función para añadir un elemento a la pila (push)

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

// Función para eliminar y devolver el elemento superior de la pila (pop)

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		panic("La pila está vacía")
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func pilas /*main*/() {

	// Crear una pila vacía

	var pila Stack

	// Añadir elementos a la pila

	pila.Push(1)
	pila.Push(2)
	pila.Push(3)

	// Eliminar y mostrar los elementos de la pila

	for len(pila) > 0 {
		fmt.Println("Elemento extraído:", pila.Pop())
	}
}

// Cola (Queue - FIFO)
// El primer elemento que se añade a la cola es el primero en ser eliminado.

// Definición de la estructura Queue

type Queue []int

// Función para añadir un elemento a la cola (enqueue)

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

// Función para eliminar y devolver el primer elemento de la cola (dequeue)

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		panic("La cola está vacía")
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func colas() {

	// Crear una cola vacía

	var cola Queue

	// Añadir elementos a la cola

	cola.Enqueue(1)
	cola.Enqueue(2)
	cola.Enqueue(3)

	// Eliminar y mostrar los elementos de la cola

	for len(cola) > 0 {
		fmt.Println("Elemento extraído:", cola.Dequeue())
	}
}