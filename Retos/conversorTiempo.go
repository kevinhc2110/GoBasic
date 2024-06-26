package main

import "fmt"

/*
Crea una función que reciba días, horas, minutos y segundos (como enteros)
y retorne su resultado en milisegundos.
*/

// Opción 1 Optima

func covertirAMilisegundos(dias, horas, minutos, segundos int) int64 {

	const (
		milisegundosPorSegundo = 1000
		segundosPorMinuto = 60
		minutosPorHora = 60
		horasPorDia = 24
	)

	totalSegundos := int64(dias*horasPorDia*minutosPorHora*segundosPorMinuto+ horas*minutosPorHora*segundosPorMinuto + minutos*segundosPorMinuto + segundos)
	totalMilisegundos := (totalSegundos * milisegundosPorSegundo)

	return totalMilisegundos
}

// Opción 2 Menos Optima

func convertirAMilisegundos1(dias, horas, minutos, segundos int) int {
	totalMilisegundos := (dias * 24 * 60 * 60 * 1000) +
		(horas * 60 * 60 * 1000) +
		(minutos * 60 * 1000) +
		(segundos * 1000)
	return totalMilisegundos
}


func convertirTiempo() {

	dias := 1
	horas := 5
	minutos := 54
	segundos := 23

	// Opción 1

	resultado := covertirAMilisegundos(dias,horas,minutos,segundos)

	fmt.Printf("El total en milisegundos es: %d\n", resultado)

	// Opción 2

	fmt.Printf("%d días, %d horas, %d minutos, %d segundos equivalen a %d milisegundos.\n",
		dias, horas, minutos, segundos, convertirAMilisegundos1(dias, horas, minutos, segundos))

} 