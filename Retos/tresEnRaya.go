package main

import "fmt"

/*
Crea una función que analice una matriz 3x3 compuesta por "X" y "O"
y retorne lo siguiente:
- "X" si han ganado las "X"
- "O" si han ganado los "O"
- "Empate" si ha habido un empate
- "Nulo" si la proporción de "X", de "O", o de la matriz no es correcta.
  O si han ganado los 2.
Nota: La matriz puede no estar totalmente cubierta.
Se podría representar con un vacío "", por ejemplo.
*/

// Función para analizar la matriz
func analizarMatriz(matriz [3][3]string) string {
	// Contadores para "X" y "O"
	countX := 0
	countO := 0

	// Variables para verificar si hay ganadores
	gananX := false
	gananO := false

	// Función interna para verificar ganador
	hayGanador := func(simbolo string) bool {
		// Verificar filas y columnas
		for i := 0; i < 3; i++ {
			if (matriz[i][0] == simbolo && matriz[i][1] == simbolo && matriz[i][2] == simbolo) ||
				(matriz[0][i] == simbolo && matriz[1][i] == simbolo && matriz[2][i] == simbolo) {
				return true
			}
		}
		// Verificar diagonales
		if (matriz[0][0] == simbolo && matriz[1][1] == simbolo && matriz[2][2] == simbolo) ||
			(matriz[0][2] == simbolo && matriz[1][1] == simbolo && matriz[2][0] == simbolo) {
			return true
		}
		return false
	}

	// Verificar el contenido de la matriz y contar símbolos al mismo tiempo
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matriz[i][j] == "X" {
				countX++
			} else if matriz[i][j] == "O" {
				countO++
			} else if matriz[i][j] != "" {
				// Si hay algún otro carácter diferente de "X", "O" o ""
				return "Nulo"
			}
		}
	}

	// Verificar ganadores después de contar los símbolos
	gananX = hayGanador("X")
	gananO = hayGanador("O")

	if gananX && gananO {
		return "Nulo" // Ambos no pueden ganar al mismo tiempo
	} else if gananX {
		return "X"
	} else if gananO {
		return "O"
	} else if countX+countO == 9 {
		return "Empate" // No hay más movimientos posibles y no hay ganador
	} else {
		return "Nulo" // Algo está mal si no se cumplen las condiciones anteriores
	}
}

func tresEnRaya() {
	// Ejemplo de uso
	matriz1 := [3][3]string{
		{"X", "O", "X"},
		{"", "X", ""},
		{"O", "O", "O"},
	}
	fmt.Println(analizarMatriz(matriz1)) // Salida esperada: "O"

	matriz2 := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "", "X"},
	}
	fmt.Println(analizarMatriz(matriz2)) // Salida esperada: "X"

	matriz3 := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "X", "O"},
	}
	fmt.Println(analizarMatriz(matriz3)) // Salida esperada: "Empate"

	matriz4 := [3][3]string{
		{"X", "X", "X"},
		{"O", "O", "O"},
		{"", "", ""},
	}
	fmt.Println(analizarMatriz(matriz4)) // Salida esperada: "Nulo"
}
