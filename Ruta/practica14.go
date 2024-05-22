package main

import (
	"fmt"
	"time"
)

/*
DIFICULTAD EXTRA:
Utilizando la fecha de tu cumpleaños, formatéala y muestra su resultado de
10 maneras diferentes. Por ejemplo:
- Día, mes y año.
- Hora, minuto y segundo.
- Día de año.
- Día de la semana.
- Nombre del mes.
(lo que se te ocurra...)
*/

func practica14() /*main*/ {
	// Crear una variable que represente tu fecha de nacimiento
	fechaNacimiento := time.Date(1997, time.October, 10, 0, 0, 0, 0, time.UTC)

	// Diferentes formatos para mostrar la fecha de nacimiento
	formatearFecha(fechaNacimiento)
}

func formatearFecha(fecha time.Time) {

	// 1. Día, mes y año

	fmt.Printf("1. Día, mes y año: %02d-%02d-%04d\n", fecha.Day(), fecha.Month(), fecha.Year())

	// 2. Hora, minuto y segundo (aunque no tenga sentido para la fecha de nacimiento, para el ejemplo)

	fmt.Printf("2. Hora, minuto y segundo: %02d:%02d:%02d\n", fecha.Hour(), fecha.Minute(), fecha.Second())

	// 3. Día del año

	fmt.Printf("3. Día del año: %03d\n", fecha.YearDay())

	// 4. Día de la semana

	fmt.Printf("4. Día de la semana: %s\n", fecha.Weekday())

	// 5. Nombre del mes

	fmt.Printf("5. Nombre del mes: %s\n", fecha.Month())

	// 6. Formato completo (RFC1123)

	fmt.Printf("6. Formato completo (RFC1123): %s\n", fecha.Format(time.RFC1123))

	// 7. Formato corto (02/01/2006)

	fmt.Printf("7. Formato corto (02/01/2006): %s\n", fecha.Format("02/01/2006"))

	// 8. Formato largo con hora (02 de enero de 2006, 15:04:05)

	fmt.Printf("8. Formato largo con hora: %s\n", fecha.Format("02 de January de 2006, 15:04:05"))

	// 9. Formato ISO 8601

	fmt.Printf("9. Formato ISO 8601: %s\n", fecha.Format("2006-01-02T15:04:05Z07:00"))

	// 10. Formato de fecha personalizada (Lunes, 02 de Enero de 2006)

	fmt.Printf("10. Formato personalizado: %s\n", fecha.Format("Monday, 02 de January de 2006"))
}


