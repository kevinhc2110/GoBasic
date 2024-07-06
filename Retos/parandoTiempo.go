package main

import (
	"fmt"
	"time"
)

/*
Crea una función que sume 2 números y retorne su resultado pasados
unos segundos.
- Recibirá por parámetros los 2 números a sumar y los segundos que
debe tardar en finalizar su ejecución.
- Si el lenguaje lo soporta, deberá retornar el resultado de forma
asíncrona, es decir, sin detener la ejecución del programa principal.
Se podría ejecutar varias veces al mismo tiempo.
*/

func sumAsync(a int, b int, delaySeconds int) <-chan int {
	// Crear un canal para transportar enteros
	result := make(chan int)

	// Lanzar una goroutine para realizar la operación asíncrona
	go func() {
		// Cerrar el canal al final de la goroutine para señalar que no se enviarán más valores
		defer close(result)

		// Esperar el tiempo especificado
		time.Sleep(time.Duration(delaySeconds) * time.Second)

		// Calcular la suma y enviarla al canal
		result <- a + b
	}()

	// Devolver el canal de solo lectura
	return result
}


func parandoTiempo() {
	a, b, delay := 5, 7, 3
	resultChan := sumAsync(a, b, delay)

	// Simulando otras operaciones mientras esperamos el resultado
	fmt.Println("Haciendo otras cosas mientras esperamos el resultado...")

	// Obtener el resultado
	result := <-resultChan
	fmt.Printf("El resultado de la suma de %d y %d es: %d\n", a, b, result)
}