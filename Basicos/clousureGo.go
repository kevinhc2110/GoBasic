package main

import "fmt"

// Un closure en Go es una función que captura variables definidas fuera de su cuerpo. Estas variables pueden ser utilizadas dentro de la función, incluso si están fuera de su alcance léxico original. Esto significa que el closure "recuerda" el estado de las variables al momento de su definición, incluso si la ejecución del programa ha avanzado más allá de ese punto.
// Captura de variables: Un closure captura las variables definidas fuera de su cuerpo.
// Estado retenido: El closure "recuerda" el estado de las variables capturadas en el momento de su definición.
// Portabilidad: Los closures pueden ser pasados como argumentos a otras funciones o ser devueltos como valores de otras funciones.
// Los closures son útiles cuando queremos crear funciones que actúen como "plantillas" y que puedan personalizarse con diferentes contextos cada vez que se llaman. Son comunes en situaciones donde necesitamos funciones que encapsulen algún comportamiento específico y que puedan ser reutilizadas en diferentes partes del código.

func clousure() /*main*/ {

	// Definimos una función llamada "incrementador" que retorna otra función.

	incrementador := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
	}

	// Creamos una instancia de la función incrementador.

	incremento := incrementador()

	// Llamamos al closure varias veces para ver cómo cambia el valor de "x".

	fmt.Println(incremento()) // Output: 1
	fmt.Println(incremento()) // Output: 2
	fmt.Println(incremento()) // Output: 3
}
