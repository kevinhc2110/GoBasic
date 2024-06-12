package main

/*
DIFICULTAD EXTRA (opcional):
Crea un decorador que sea capaz de contabilizar cuántas veces
se ha llamado a una función y aplícalo a una función de tu elección.
*/

import (
	"fmt"
	"sync"
)

// Contador struct para mantener el conteo de llamadas
type Contador struct {
	mu       sync.Mutex
	funcion  func()
	numLlamadas int
}

// Es una función constructora que crea y devuelve una nueva instancia del tipo Contador.
// NewContador crea un nuevo Contador para una función dada
func NewContador(f func()) *Contador {
	// f es una función que no recibe argumentos y no retorna nada (func()). Este es el tipo de la función que queremos decorar.
	// La función retorna un puntero a una nueva instancia de Contador (*Contador).
	return &Contador{funcion: f}
	// Inicializa el campo funcion de esta nueva instancia con la función f que se pasa como argumento.
}

// Decorator método que cuenta las llamadas y luego ejecuta la función original
func (c *Contador) Decorator() {
	c.mu.Lock()
	c.numLlamadas++
	c.mu.Unlock()
	c.funcion()
}

// GetCount devuelve el número de veces que se ha llamado a la función
func (c *Contador) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.numLlamadas
}

// Función de prueba a la que se le aplicará el decorador
func miFuncion() {
	fmt.Println("Ejecutando la función original")
}

func practica24() {
	// Crear un contador para miFuncion
	contador := NewContador(miFuncion)

	// Llamar a la función decorada varias veces
	contador.Decorator()
	contador.Decorator()
	contador.Decorator()

	// Imprimir el número de veces que se ha llamado a la función
	fmt.Printf("La función se ha llamado %d veces\n", contador.GetCount())
}