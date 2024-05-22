package main

import (
	"errors"
	"fmt"
)

/*
DIFICULTAD EXTRA:
Implementa dos clases que representen las estructuras de Pila y Cola (estudiadas
en el ejercicio número 7 de la ruta de estudio)
- Deben poder inicializarse y disponer de operaciones para añadir, eliminar,
retornar el número de elementos e imprimir todo su contenido.
*/

// Definición de la estructura de la Pila

type Stack1 struct {
	data []int
}

// Método para agregar un elemento a la Pila

func (stack *Stack1) Push(element int) {
	stack.data = append(stack.data, element)
}

// Método para eliminar y devolver el último elemento de la Pila

func (stack *Stack1) Pop() (int, error) {
	if len(stack.data) == 0 {
		return 0, errors.New("pila vacía")
	}
	element := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return element, nil
}

// Método para retornar el número de elementos en la Pila

func (stack *Stack1) Size() int {
	return len(stack.data)
}

// Método para imprimir todo el contenido de la Pila

func (stack *Stack1) Print() {
	for i := len(stack.data) - 1; i >= 0; i-- {
		fmt.Println("Elemento:", stack.data[i])
	}
}

// Definición de la estructura de la Cola

type Queue1 struct {
	data []int
}

// Método para agregar un elemento a la Cola

func (queue *Queue1) Enqueue(element int) {
	queue.data = append(queue.data, element)
}

// Método para eliminar y devolver el primer elemento de la Cola

func (queue *Queue1) Dequeue() (int, error) {
	if len(queue.data) == 0 {
		return 0, errors.New("cola vacía")
	}
	element := queue.data[0]
	queue.data = queue.data[1:]
	return element, nil
}

// Método para retornar el número de elementos en la Cola

func (queue *Queue1) Size() int {
	return len(queue.data)
}

// Método para imprimir todo el contenido de la Cola

func (queue *Queue1) Print() {
	for _, element := range queue.data {
		fmt.Println("Elemento:", element)
	}
}

func practica08()/*main*/ {

	// Crear una pila y agregar elementos

	stack := Stack1{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Imprimir contenido de la pila

	fmt.Println("Contenido de la pila:")
	stack.Print()
	fmt.Println("Número de elementos en la pila:", stack.Size())

	// Crear una cola y agregar elementos

	queue := Queue1{}
	queue.Enqueue(100)
	queue.Enqueue(200)
	queue.Enqueue(300)

	// Imprimir contenido de la cola

	fmt.Println("\nContenido de la cola:")
	queue.Print()
	fmt.Println("Número de elementos en la cola:", queue.Size())
}