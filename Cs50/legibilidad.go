package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

//* Problema 2 Readability

func legibilidad() {

	fmt.Println("Ingrese el texto a analizar")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	indice := indiceColemanLiai(texto)

	if (indice >= 16){

			fmt.Println("Grado 16+");

	}else if (indice < 1){

		fmt.Println("Antes de grado 1");
		
	}else{
		fmt.Printf("Grado %d\n", indice);
	}
	
}

func indiceColemanLiai(texto string) int {

	/*
	for _, letra := range texto {
    if unicode.IsLetter(letra){
      conteoLetras++
    }
  }
	*/

	// En lugar de iterar sobre las palabras (strings.Fields(texto)), se podría iterar sobre caracteres individuales (rune) utilizando un bucle for range y verificar si cada carácter es una letra usando unicode.IsLetter(rune).
	// Eficiencia 

	conteoLetras := 0

	for i := 0; i < len(texto); i++ {
		letra := texto[i]
		if unicode.IsLetter(rune(letra)) {
			conteoLetras++
		}
	}

	palabras := strings.Fields(texto)
	conteoPalabras := len(palabras)

	conteoOraciones := 0
	
	for _, palabra := range palabras {
		if strings.ContainsAny(palabra, ".!?") {
			conteoOraciones++
		}
	}

	l := (float64(conteoLetras) / float64(conteoPalabras)) * 100
	s := (float64(conteoOraciones) / float64(conteoPalabras)) * 100
	indice := (0.0588*l - 0.296*s - 15.8)

	return int(math.Round(indice))
} 





