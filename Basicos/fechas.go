package main

import (
	"fmt"
	"time"
)

func fechas() /*main*/ {

	// Crear una variable que represente la fecha actual
	// time.Now() retorna la fecha y hora actuales del sistema.

	fechaActual := time.Now()
	fmt.Printf("Fecha actual: %v\n", fechaActual)

	// Crear una variable que represente la fecha de nacimiento
	// time.Date() crea una nueva fecha a partir de los valores proporcionados: año, mes, día, hora, minuto, segundo, nanosegundo y la zona horaria.
	// 1990: Año.
	// time.January: Mes (Enero).
	// 1: Día.
	// 12, 0, 0, 0: Hora (12:00:00:000).
	// time.UTC: Zona horaria (UTC).

	fechaNacimiento := time.Date(1990, time.January, 1, 12, 0, 0, 0, time.UTC)
	fmt.Printf("Fecha de nacimiento: %v\n", fechaNacimiento)

	// Calcular la diferencia en años

	edad := calcularAnios(fechaNacimiento, fechaActual)
	fmt.Printf("Han transcurrido %d años desde la fecha de nacimiento.\n", edad)

		// Extraer y mostrar la hora actual

		horaActual := fechaActual.Hour()
		minutoActual := fechaActual.Minute()
		segundoActual := fechaActual.Second()

		fmt.Printf("Hora actual: %02d:%02d:%02d\n", horaActual, minutoActual, segundoActual)
}

// calcularAnios calcula la diferencia en años entre dos fechas
// Year() retorna el año de la fecha especificada
// YearDay() retorna el día del año (de 1 a 365 o 366 en años bisiestos) de la fecha especificada.

func calcularAnios(fechaInicio, fechaFin time.Time) int {
	years := fechaFin.Year() - fechaInicio.Year()
	if fechaFin.YearDay() < fechaInicio.YearDay() {
		years--
	}
	return years
}