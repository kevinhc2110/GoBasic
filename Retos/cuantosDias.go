package main

import (
	"fmt"
	"time"
)

/*
Crea una función que calcule y retorne cuántos días hay entre dos cadenas
de texto que representen fechas.
- Una cadena de texto que representa una fecha tiene el formato "dd/MM/yyyy".
- La función recibirá dos String y retornará un Int.
- La diferencia en días será absoluta (no importa el orden de las fechas).
- Si una de las dos cadenas de texto no representa una fecha correcta se
  lanzará una excepción.
*/

func cuantosDias() {

	fecha1 := "10/06/2024"
	fecha2 := "05/06/2024"

	dias, err := diferenciaDias(fecha1, fecha2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Diferencia en días entre %s y %s: %d días\n", fecha1, fecha2, dias)
}

func diferenciaDias(fecha1 string, fecha2 string) (int, error) {

	// Analizar las cadenas de texto y convertirlas en objetos time.Time
	// El primer argumento ("02/01/2006") especifica el diseño esperado de la cadena de fecha

	fechaTime1, err := time.Parse("02/01/2006", fecha1)
	if err != nil {
		// Se retorna 0 en el caso de que una fecha no sea válida es para indicar claramente que no se pudo calcular la diferencia entre las fechas debido a un problema con una de ellas
		return 0, fmt.Errorf("fecha1 no es válida: %v", err)
	}

	fechaTime2, err := time.Parse("02/01/2006", fecha2)
	if err != nil {
		return 0, fmt.Errorf("fecha2 no es válida: %v", err)
	}

	// Calcular la diferencia en días entre las dos fechas
	// La expresión fechaTime1.Sub(fechaTime2).Hours() / 24 está calculando la diferencia en horas entre fechaTime1 y fechaTime2, y luego la convierte en días dividiéndola por 24 (ya que hay 24 horas en un día).

	diferencia := fechaTime1.Sub(fechaTime2).Hours() / 24
	if diferencia < 0 {
		// Para asegurar que sean siempre valores positivos if diferencia < 0 para cambiar el signo de la diferencia a positivo multiplicándola por -1
		diferencia = -diferencia
	}

	return int(diferencia), nil

}
