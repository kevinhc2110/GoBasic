package main

import (
	"fmt"
	"strings"
)

func carreraObstaculos() {

	acciones := []string{"run", "jump", "run", "run", "jump"}
	pista := "_|__|"
	resultado := evaluarCarrera(acciones, pista)
	resultado2 := evaluarCarrera2(acciones, pista)

	// Opción 1

	if resultado {
		fmt.Println("El/a atleta ha superado la carrera correctamente.")
	} else {
		fmt.Println("El/a atleta no ha superado la carrera.")
	}

	// Opción 2

	if resultado2 {
		fmt.Println("El/a atleta ha superado la carrera correctamente.")
	} else {
		fmt.Println("El/a atleta no ha superado la carrera.")
	}

}

// Opción 1

func evaluarCarrera(acciones []string, pista string) bool {

	// Convertir la pista en un slice de runes para poder modificarla

	pistaRunes := []rune(pista)
	superado := true

	// Asegurarse de que la longitud de acciones y pista sea la misma

	if len(acciones) != len(pistaRunes) {
		fmt.Println("La longitud de las acciones y la pista no coinciden.")
		return false
	}

	// Iterar sobre las acciones y la pista

	for i, accion := range acciones {

		if (accion == "run" && pistaRunes[i] == '_') || (accion == "jump" && pistaRunes[i] == '|') {
			// Acción correcta, no hacer nada
		} else if accion == "run" && pistaRunes[i] == '|'  {
			// Jump en suelo, marcar como incorrecto
			pistaRunes[i] = '/'
			return false
		} else if accion == "jump" && pistaRunes[i] == '_' {
			// Run en valla, marcar como incorrecto
			pistaRunes[i] = 'x'
			return false
		}		
	}

	// Imprimir el resultado de la pista
	fmt.Println(string(pistaRunes))


	return superado

}

// Opción 2

func evaluarCarrera2(acciones []string, pista string) bool {
	if len(acciones) != len(pista) {
		fmt.Println("La longitud de las acciones y la pista no coinciden.")
		return false
	}

	var pistaFinal strings.Builder
	superado := true

	for i := 0; i < len(acciones); i++ {
		accion := acciones[i]
		tramo := pista[i]

		switch {
		case accion == "run" && tramo == '_':
			pistaFinal.WriteRune('_')
		case accion == "jump" && tramo == '|':
			pistaFinal.WriteRune('|')
		case accion == "jump" && tramo == '_':
			pistaFinal.WriteRune('x')
			superado = false
		case accion == "run" && tramo == '|':
			pistaFinal.WriteRune('/')
			superado = false
		default:
			// En caso de acción desconocida, marcar como incorrecto
			pistaFinal.WriteRune('?')
			superado = false
		}
	}

	fmt.Println(pistaFinal.String())
	return superado
}