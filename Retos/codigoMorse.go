package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

func main() {

	fmt.Println("Ingrese un texto o un código morse para convertir")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	tipoCoversion := detectarTipo(texto)

	if tipoCoversion == "texto" {
		println(conversionMorse(texto))
	}else if tipoCoversion == "morse"{
		println(conversionTexto(texto))
	}else {
		fmt.Println("Entrada no validad, intente de nuevo")
	}

}

func conversionMorse(texto string) string {

	textoMayuscula := strings.ToUpper(texto)
	conversionMorse := ""

	for _, codigo := range textoMayuscula {

		if codigo >= 'A' && codigo <= 'Z' || codigo >= '0' && codigo <= '9' || codigo == '.' || codigo == ',' || codigo == '?' || codigo == '!'  {
			conversionMorse += diccionarioMorse[string(codigo)] + " "
		} else {
			conversionMorse += string(codigo) 
		}
	}
	conversionMorse = strings.TrimSpace(conversionMorse)
	return conversionMorse
}

func conversionTexto(morse string) string {
	
	// Reemplaza todas las ocurrencias de doble espacio ("  ") por un solo espacio (" ")
	morseSinEspacios := strings.ReplaceAll(morse, "  ", " ")
	// Divide la cadena de Morse en palabras separadas por un espacio
	palabras := strings.Fields(morseSinEspacios)
	textoDecodificado := ""

	// Itera sobre cada palabra en la lista de palabras Morse
	for _, palabra := range palabras {

			// Divide la palabra Morse en letras individuales separadas por un espacio
			letras := strings.Fields(palabra)
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
			textoDecodificado += " "
	}
	textoDecodificado = strings.TrimSpace(textoDecodificado)
	return textoDecodificado
}

func detectarTipo(entrada string) string {

	if regexp.MustCompile("[A-Z]").MatchString(entrada) {
			return "texto"
	} else if regexp.MustCompile("[.-]").MatchString(entrada) {
			return "morse"
	} else {
			return "invalido"
	}
}