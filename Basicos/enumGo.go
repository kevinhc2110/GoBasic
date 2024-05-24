package main

import "fmt"

// Es un tipo de dato especial que consiste en un conjunto de valores constantes, llamados "enumeradores", que representan un conjunto finito de posibles valores para una variable.
// En Go, no hay un tipo de dato específico llamado enum como en algunos otros lenguajes de programación. Sin embargo, puedes simular enumeraciones utilizando constantes sin tipo y crear funciones o métodos asociados para trabajar con ellas.

// Definir un tipo para los días de la semana

type DiaSemana int

// Definir constantes para los días de la semana

const (

	// Empezamos desde 1 utilizando iota, que es una constante predeclarada que genera valores secuenciales.

	Lunes DiaSemana = iota + 1
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)

// Función para obtener el nombre del día de la semana

func NombreDia(dia DiaSemana) string {
	switch dia {
	case Lunes:
		return "Lunes"
	case Martes:
		return "Martes"
	case Miercoles:
		return "Miércoles"
	case Jueves:
		return "Jueves"
	case Viernes:
		return "Viernes"
	case Sabado:
		return "Sábado"
	case Domingo:
		return "Domingo"
	default:
		return "Día no válido"
	}
}

func enumGo() /*main*/ {
	// Ejemplo de uso de la función dayToString para el número 1
	dia := 1
	fmt.Printf("Día %d: %s\n", dia, NombreDia(DiaSemana(dia)))

	for i := 1; i <= 7; i++ {
		fmt.Printf("Día %d: %s\n", i, NombreDia(DiaSemana(i)))
	}
}