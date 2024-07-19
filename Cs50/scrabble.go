package main

import (
	"fmt"
	"strings"
)

//* Problema 2 Scrabble

func scrabbleValues(palabra string) int {

	scrabbleValues := map[rune]int{
		'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1, 'F': 4, 'G': 2, 'H': 4,
		'I': 1, 'J': 8, 'K': 5, 'L': 1, 'M': 3, 'N': 1, 'O': 1, 'P': 3,
		'Q': 10, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 4, 'W': 4, 'X': 8,
		'Y': 4, 'Z': 10,
	}

	puntaje := 0

	for _, valor := range palabra {
		puntaje += scrabbleValues[valor]
	}
	return puntaje
}

func scrabble() {

	var palabra1 string
	var palabra2 string

	fmt.Print("Player 1: ")
	fmt.Scanln(&palabra1)

	fmt.Print("Player 2: ")
	fmt.Scanln(&palabra2)

	palabra1 = strings.ToUpper(palabra1)
	palabra2 = strings.ToUpper(palabra2)

	if scrabbleValues(palabra1) > scrabbleValues(palabra2) {
		fmt.Println("Player 1: Wins!")
	} else if scrabbleValues(palabra1) == scrabbleValues(palabra2) {
		fmt.Println("Tie!")
	} else {
		fmt.Println("Player 2: Wins!")
	}
}
