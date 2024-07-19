//* Callbacks

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* DIFICULTAD EXTRA:
Crea un simulador de pedidos de un restaurante utilizando callbacks.
Estará formado por una función que procesa pedidos.
Debe aceptar el nombre del plato, una callback de confirmación, una
de listo y otra de entrega.
- Debe imprimir un confirmación cuando empiece el procesamiento.
- Debe simular un tiempo aleatorio entre 1 a 10 segundos entre
  procesos.
- Debe invocar a cada callback siguiendo un orden de procesado.
- Debe notificar que el plato está listo o ha sido entregado.
*/

// Definición de las funciones de callback

type Callback func(string)

func practica21() /*main*/ {

	// Función principal para procesar pedidos

	procesarPedido("Hamburguesa", confirmarPedido, platoListo, platoEntregado)
}

// Función para procesar pedidos

func procesarPedido(plato string, confirmar Callback, listo Callback, entregado Callback) {

	// Imprimir confirmación de inicio de procesamiento

	fmt.Println("Iniciando procesamiento del pedido de", plato)

	// Llamar a la función de confirmación

	confirmar(plato)

	// Simular un tiempo aleatorio entre 1 y 10 segundos

	tiempoAleatorio := rand.Intn(10) + 1
	time.Sleep(time.Duration(tiempoAleatorio) * time.Second)

	// Llamar a la función de plato listo

	listo(plato)

	// Simular otro tiempo aleatorio entre 1 y 10 segundos

	tiempoAleatorio = rand.Intn(10) + 1
	time.Sleep(time.Duration(tiempoAleatorio) * time.Second)

	// Llamar a la función de plato entregado

	entregado(plato)
}

// Función de callback para confirmar pedido

func confirmarPedido(plato string) {
	fmt.Println("Pedido de", plato, "confirmado.")
}

// Función de callback para plato listo

func platoListo(plato string) {
	fmt.Println("El plato de", plato, "está listo para servir.")
}

// Función de callback para plato entregado

func platoEntregado(plato string) {
	fmt.Println("El plato de", plato, "ha sido entregado al cliente.")
}
