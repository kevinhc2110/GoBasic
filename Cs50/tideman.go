package main

import (
	"fmt"
	"os"
	"sort"
)

//* Problema 2 Tideman

// preferences[i][j] es el número de votantes que prefieren i sobre j
var preferences [MAX][MAX]int

// locked[i][j] significa que i está bloqueado sobre j
var locked [MAX][MAX]bool

// Cada par tiene un ganador y un perdedor
type pair struct {
	winner int
	loser  int
}

// Array de candidatos
var candidates [MAX]string
var pairs []pair

var pairCount int
var candidateCount int

func tideman() {
	// Verificar el uso inválido
	if len(os.Args) < 2 {
		fmt.Println("Uso: tideman [candidato ...]")
		os.Exit(1)
	}

	// Poblar el array de candidatos
	candidateCount = len(os.Args) - 1
	if candidateCount > MAX {
		fmt.Printf("El número máximo de candidatos es %d\n", MAX)
		os.Exit(2)
	}
	for i := 0; i < candidateCount; i++ {
		candidates[i] = os.Args[i+1]
	}

	// Limpiar el gráfico de pares bloqueados
	for i := 0; i < candidateCount; i++ {
		for j := 0; j < candidateCount; j++ {
			locked[i][j] = false
		}
	}

	pairCount = 0
	numeroDeVotantes := ObtenerEntero("Número de votantes: ")

	// Consultar los votos
	for i := 0; i < numeroDeVotantes; i++ {
		// ranks[i] es la preferencia i del votante
		ranks := make([]int, candidateCount)

		// Consultar cada rango
		for j := 0; j < candidateCount; j++ {
			nombre := ObtenerCadena(fmt.Sprintf("Rank %d: ", j+1))

			if !vote(j, nombre, ranks) {
				fmt.Println("Voto inválido.")
				os.Exit(3)
			}
		}

		recordPreferences(ranks)
		fmt.Println()
	}

	addPairs()
	sortPairs()
	lockPairs()
	printWinner()

}

// Actualizar los rangos dado un nuevo voto
func vote(rank int, name string, ranks []int) bool {
	for i := 0; i < candidateCount; i++ {
		if candidates[i] == name {
			ranks[rank] = i
			return true
		}
	}
	return false
}

// Actualizar las preferencias dado los rangos de un votante
// Cada fila y columna representa un candidato.
// El valor en preferences[i][j] indica cuántos votantes prefieren al candidato i sobre el candidato j

func recordPreferences(ranks []int) {
	for i := 0; i < candidateCount; i++ {
		for j := i + 1; j < candidateCount; j++ {
			preferences[ranks[i]][ranks[j]]++
		}
	}
}

// Registrar pares de candidatos donde uno es preferido sobre el otro
func addPairs() {
	for i := 0; i < candidateCount; i++ {
		for j := 0; j < candidateCount; j++ {
			if preferences[i][j] > preferences[j][i] {
				pairs = append(pairs, pair{winner: i, loser: j})
			}
		}
	}
	pairCount = len(pairs)
}

// Ordenar los pares en orden decreciente por la fuerza de la victoria
func sortPairs() {
	sort.Slice(pairs, func(i, j int) bool {
		return preferences[pairs[i].winner][pairs[i].loser] > preferences[pairs[j].winner][pairs[j].loser]
	})
}

// Bloquear pares en el gráfico de candidatos en orden, sin crear ciclos
func lockPairs() {
	for i := 0; i < pairCount; i++ {
		if !createsCycle(pairs[i].winner, pairs[i].loser) {
			locked[pairs[i].winner][pairs[i].loser] = true
		}
	}
}

// Verificar si bloquear un par crea un ciclo
func createsCycle(winner, loser int) bool {
	if winner == loser {
		return true
	}
	for i := 0; i < candidateCount; i++ {
		if locked[loser][i] && createsCycle(winner, i) {
			return true
		}
	}
	return false
}

// Imprimir el ganador de la elección
func printWinner() {
	for i := 0; i < candidateCount; i++ {
		isWinner := true
		for j := 0; j < candidateCount; j++ {
			if locked[j][i] {
				isWinner = false
				break
			}
		}
		if isWinner {
			fmt.Println(candidates[i])
			return
		}
	}
}
