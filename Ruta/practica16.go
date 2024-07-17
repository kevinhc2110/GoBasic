//* Expresiones regulares

package main

import (
	"fmt"
	"regexp"
)

/*
DIFICULTAD EXTRA:
Crea 3 expresiones regulares (a tu criterio) capaces de:
- Validar un email.
- Validar un número de teléfono.
- Validar una url.
*/

func practica16() /*main*/ {

	email := "usuario@example.com"
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if match {
			fmt.Println("El email es válido")
	} else {
			fmt.Println("El email no es válido")
	}

	telefono := "+1 (123) 456-7890"
	match1, _ := regexp.MatchString(`^\+\d{1,3}\s?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`, telefono)
	if match1 {
			fmt.Println("El número de teléfono es válido")
	} else {
			fmt.Println("El número de teléfono no es válido")
	}

	url := "https://www.example.com"
	match2, _ := regexp.MatchString(`^(https?|ftp):\/\/[^\s/$?#].[^\s\/]*\.[a-zA-Z]{2,}(\/|\?|$)`, url)
	if match2 {
			fmt.Println("La URL es válida")
	} else {
			fmt.Println("La URL no es válida")
	}

}