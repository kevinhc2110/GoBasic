package main

import (
	"fmt"
	"strings"
)

/*
Crea una función que reciba dos cadenas como parámetro (str1, str2)
e imprima otras dos cadenas como salida (out1, out2).
- out1 contendrá todos los caracteres presentes en la str1 pero NO
  estén presentes en str2.
- out2 contendrá todos los caracteres presentes en la str2 pero NO
  estén presentes en str1.
*/

func eliminandoCaracteres()  {

	str1 := "abcde"
	str2 := "cdefg"

	// Opción 1

	out1, out2 := convertirPalabras1(str1, str2)
	fmt.Printf("Caracteres únicos de str1: %s\n", out1)
	fmt.Printf("Caracteres únicos de str2: %s\n", out2)

	// Opción 2
	
	out1, out2 = convertirPalabras2(str1, str2)
	fmt.Printf("Caracteres únicos de str1: %s\n", out1)
	fmt.Printf("Caracteres únicos de str2: %s\n", out2)
	
}

// Opción 1

func convertirPalabras1(str1 string, str2 string)(out1 string, out2 string)  {

	// Uso de strings.Builder para construir las cadenas de salida
	// Las variables deben ser nombradas con builder

	var builderOut1 strings.Builder
	var builderOut2 strings.Builder

	// Mapa para verificar la existencia de caracteres en str2
	
	caracteresStr2 := make(map[rune]bool)

	for _, char := range str2 {
		caracteresStr2[char] = true
	}

	// Agregar caracteres únicos de str1 a out1

	for _, char := range str1 {
		if !caracteresStr2[char] {

			builderOut1.WriteRune(char)
			
		}
	}
	
	// Mapa para verificar la existencia de caracteres en str1
	
	caracteresStr1 := make(map[rune]bool)

	for _, char := range str1 {
		caracteresStr1[char] = true
	}

	// Agregar caracteres únicos de str2 a out2

	for _, char := range str2 {
		if !caracteresStr1[char] {

			builderOut2.WriteRune(char)
			
		}
	}

	return builderOut1.String(), builderOut2.String()
}

// Opción 2

func convertirPalabras2(str1 string, str2 string) (out1 string, out2 string) {

	set1 := make(map[rune]bool) // Crea un conjunto para caracteres str1
	for _, char := range str1 {
			set1[char] = true
	}

	out1, out2 = "", ""

	for _, char := range str2 {
			if !set1[char] { // Si char no está presente en set1 (único en str2)
					out2 += string(char)
			} else {
					delete(set1, char) // Elimina char de set1 si se encuentra en str2
					if len(set1) == 0 {
							break // Terminación anticipada si todos los caracteres en str1 se encuentran en str2
					}
			}
	}

	for char := range set1 { // Agrega los caracteres únicos restantes de str1
			out1 += string(char)
	}

	return out1, out2
}