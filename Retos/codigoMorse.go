package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

/*
Crea un programa que sea capaz de transformar texto natural a código
morse y viceversa.
- Debe detectar automáticamente de qué tipo se trata y realizar
  la conversión.
- En morse se soporta raya "—", punto ".", un espacio " " entre letras
  o símbolos y dos espacios entre palabras "  ".
- El alfabeto morse soportado será el mostrado en
  https://es.wikipedia.org/wiki/Código_morse.
*/

// Aunque es una opción es ams eficiente usar claves como runas ' '

var diccionarioMorse = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	".": ".-.-.-.",
	",": "--..--",
	"?": "..--..",
	"!": "-.-.--.-",
}

func codigoMorse() {

	fmt.Println("Ingrese un texto o un código morse para convertir")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	textoMayuscula := strings.ToUpper(texto)

	// Opcion 1 DetectatTipo

	tipoConversion := detectarTipo(textoMayuscula)

	if tipoConversion == "texto" {
		fmt.Println(conversionMorse(textoMayuscula))
	}else if tipoConversion == "morse"{
		fmt.Println(conversionTexto(textoMayuscula))
	}else {
		fmt.Println("Entrada no validad, intente de nuevo")
	}

	// Opcion 2 DetectatTipo

	tipoConversion2 := detectarTipo(textoMayuscula)

	if tipoConversion2 == "texto" {
		fmt.Println(conversionMorse(textoMayuscula))
	}else if tipoConversion2 == "morse"{
		fmt.Println(conversionTexto(textoMayuscula))
	}else {
		fmt.Println("Entrada no validad, intente de nuevo")
	}

}

func conversionMorse(texto string) string {

	conversionMorse := ""

	for _, codigo := range texto {

		if codigo >= 'A' && codigo <= 'Z' || codigo >= '0' && codigo <= '9' || codigo == '.' || codigo == ',' || codigo == '?' || codigo == '!'  {
			conversionMorse += diccionarioMorse[string(codigo)] + " "
		} else {
			// Si el carácter no cumple la condición se agrega directamente
			conversionMorse += string(codigo) 
		}
	}
	conversionMorse = strings.TrimSpace(conversionMorse)
	return conversionMorse
}

func conversionTexto(morse string) string {
	
	// Divide la cadena de Morse en palabras separadas por dos espacios, se genera un array
	palabras := strings.Split(morse, "  ")
	
	textoDecodificado := ""

	// Itera sobre cada palabra en la lista de palabras Morse
	for _, palabra := range palabras {
		// Divide la palabra Morse en letras individuales separadas por un espacio
		letras := strings.Split(palabra, " ")
		// Itera sobre cada letra en la palabra Morse
		for _, letra := range letras {
			// Busca en el diccionario Morse la clave correspondiente al valor (letra Morse)
			for clave, valor := range diccionarioMorse {
				if valor == letra {
					textoDecodificado += clave
					break
				}
			}
		}
		// Esto genera espacio en el principio, en la separación y en final
		textoDecodificado += " "
	}
	// Con esto eliminamos los espacio en el inicio y en el final
	textoDecodificado = strings.TrimSpace(textoDecodificado)
	return textoDecodificado
}

// Opción 1 DetectatTipo

func detectarTipo(entrada string) string {

	if regexp.MustCompile("[A-Z0-9.,?!-]").MatchString(entrada) {
			return "texto"
	} else if regexp.MustCompile("[.-]").MatchString(entrada) {
			return "morse"
	} else {
			return "invalido"
	}
}

// Opcion 2 DetectatTipo

func detectarTipo2(entrada string) string {

	esTexto := false
	esMorse := false

	for _, char := range entrada {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || char == '.' || char == ',' || char == '?' || char == '!' || char == '-' {
			esTexto = true
		} else if char == '.' || char == '-' {
			esMorse = true
		} else {
			// Si encontramos un caracter que no es ni texto ni morse, retornamos "invalido"
			return "invalido"
		}
	}

	if esTexto && !esMorse {
		return "texto"
	} else if esMorse && !esTexto {
		return "morse"
	} else {
		// Si la cadena contiene caracteres de ambos tipos, retornamos "invalido"
		return "invalido"
	}
}