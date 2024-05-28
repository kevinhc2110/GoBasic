package main

import (
	"fmt"
	"time"
)

// Función que se ejecutará de manera asíncrona
// Recibe el nombre de la tarea, la duración en segundos y un canal done para notificar cuando termina.
// Simula la duración de la ejecución usando time.Sleep.
// Imprime el tiempo de finalización y envía un valor true al canal done para indicar que ha terminado.

func ejecutarAsincrono(nombre string, duracion int, done chan bool) {
	fmt.Printf("La función '%s' empieza a las %v\n", nombre, time.Now())
	fmt.Printf("La función '%s' se ejecutará durante %d segundos\n", nombre, duracion)
	time.Sleep(time.Duration(duracion) * time.Second)
	fmt.Printf("La función '%s' terminó a las %v\n", nombre, time.Now())
	done <- true
}

func asincroniaGoroutines() {

	// Crear un canal para sincronización

	done := make(chan bool)

	// Ejecutar la  función en una Goroutine para que se ejecute de manera asíncrona

	go ejecutarAsincrono("Tarea Asíncrona", 5, done)

	// Espera a que la función termine leyendo del canal done.
	// El canal done es un mecanismo de sincronización utilizado para coordinar y esperar la finalización de las Goroutines (funciones concurrentes) 

	<-done
	fmt.Println("La ejecución ha finalizado")
}

