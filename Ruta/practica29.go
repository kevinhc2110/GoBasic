//* Solid: Principio de segregación de interfaces (ISP)

package main

import (
	"fmt"
)

/* DIFICULTAD EXTRA:
Crea un gestor de impresoras.
Requisitos:
1. Algunas impresoras sólo imprimen en blanco y negro.
2. Otras sólo a color.
3. Otras son multifunción, pueden imprimir, escanear y enviar fax.
Instrucciones:
1. Implementa el sistema, con los diferentes tipos de impresoras y funciones.
2. Aplica el ISP a la implementación.
3. Desarrolla un código que compruebe que se cumple el principio.
*/

// Interfaz para imprimir en blanco y negro
type ImpresoraBlancoYNegro interface {
	ImprimirBlancoYNegro()
}

// Interfaz para imprimir a color
type ImpresoraColor interface {
	ImprimirColor()
}

// Interfaz para escanear
type Escaner interface {
	Escanear()
}

// Interfaz para enviar fax
type Fax interface {
	EnviarFax()
}

// ImpresoraBlancoYNegroImpl es una impresora que solo imprime en blanco y negro
type ImpresoraBlancoYNegroImpl struct{}

func (i ImpresoraBlancoYNegroImpl) ImprimirBlancoYNegro() {
	fmt.Println("Imprimiendo en blanco y negro...")
}

// ImpresoraColorImpl es una impresora que solo imprime a color
type ImpresoraColorImpl struct{}

func (i ImpresoraColorImpl) ImprimirColor() {
	fmt.Println("Imprimiendo a color...")
}

// ImpresoraMultifuncion es una impresora que puede imprimir, escanear y enviar fax
type ImpresoraMultifuncion struct{}

func (i ImpresoraMultifuncion) ImprimirBlancoYNegro() {
	fmt.Println("Imprimiendo en blanco y negro...")
}

func (i ImpresoraMultifuncion) ImprimirColor() {
	fmt.Println("Imprimiendo a color...")
}

func (i ImpresoraMultifuncion) Escanear() {
	fmt.Println("Escaneando documento...")
}

func (i ImpresoraMultifuncion) EnviarFax() {
	fmt.Println("Enviando fax...")
}

func practica29() {
	var impresoraBN ImpresoraBlancoYNegro = ImpresoraBlancoYNegroImpl{}
	var impresoraColor ImpresoraColor = ImpresoraColorImpl{}
	var impresoraMulti ImpresoraMultifuncion = ImpresoraMultifuncion{}

	// Usar la impresora en blanco y negro
	impresoraBN.ImprimirBlancoYNegro()

	// Usar la impresora a color
	impresoraColor.ImprimirColor()

	// Usar la impresora multifunción
	impresoraMulti.ImprimirBlancoYNegro()
	impresoraMulti.ImprimirColor()
	impresoraMulti.Escanear()
	impresoraMulti.EnviarFax()
}
