//* Asincronías

package main

import (
	"fmt"
	"time"
)

/*
DIFICULTAD EXTRA:
Utilizando el concepto de asincronía y la función anterior, crea
el siguiente programa que ejecuta en este orden:
- Una función C que dura 3 segundos.
- Una función B que dura 2 segundos.
- Una función A que dura 1 segundo.
- Una función D que dura 1 segundo.
- Las funciones C, B y A se ejecutan en paralelo.
- La función D comienza su ejecución cuando las 3 anteriores han finalizado.
*/

// Función que se ejecutará de manera asíncrona

func ejecutarAsincrono(nombre string, duracion int, done chan bool) {

	fmt.Printf("La función '%s' empieza a las %v\n", nombre, time.Now())
	fmt.Printf("La función '%s' se ejecutará durante %d segundos\n", nombre, duracion)
	time.Sleep(time.Duration(duracion) * time.Second)
	fmt.Printf("La función '%s' terminó a las %v\n", nombre, time.Now())
	done <- true
}

// Función A

func funcionA(done chan bool) {
	ejecutarAsincrono("Función A", 1, done)
}

// Función B

func funcionB(done chan bool) {
	ejecutarAsincrono("Función B", 2, done)
}

// Función C

func funcionC(done chan bool) {
	ejecutarAsincrono("Función C", 3, done)
}

// Función D

func funcionD(done chan bool) {
	ejecutarAsincrono("Función D", 1, done)
}

func practica15() /*main*/ {

	// Crear canales para sincronización

	doneA := make(chan bool)
	doneB := make(chan bool)
	doneC := make(chan bool)
	doneD := make(chan bool)

	// Ejecutar las funciones A, B y C en paralelo

	go funcionA(doneA)
	go funcionB(doneB)
	go funcionC(doneC)

	// Esperar a que A, B y C terminen

	<-doneA
	<-doneB
	<-doneC

	// Ejecutar la función D después de que A, B y C hayan terminado

	go funcionD(doneD)

	// Esperar a que D termine

	<-doneD
	fmt.Println("Todas las funciones han finalizado")
}
