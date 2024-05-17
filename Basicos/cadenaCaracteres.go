package main

import (
	"fmt"
	"strings"
)

func cadenaCaracteres /*main*/ () {

	// Acceso a caracteres específicos

	cadena := "Hola Mundo"
	fmt.Println("Carácter en la posición 4:", cadena[3]) // Imprime: 97 (código ASCII de 'a')

	// Subcadenas

	fmt.Println("Subcadena desde la posición 5:", cadena[4:]) // Imprime: "Mundo"
	fmt.Println("Subcadena hasta la posición 3:", cadena[:3]) // Imprime: "Hol"

	// Longitud

	fmt.Println("Longitud de la cadena:", len(cadena)) // Imprime: 10

	// Concatenación

	cadena1 := "Hola"
	cadena2 := "Mundo"
	concatenada := cadena1 + " " + cadena2
	fmt.Println("Cadena concatenada:", concatenada) // Imprime: "Hola Mundo"

	// Repetición

	cadena3 := strings.Repeat("abc", 3)
	fmt.Println("Cadena repetida:", cadena3) // Imprime: "abcabcabc"

	// Recorrido

	for _, char := range cadena {
		fmt.Println("Carácter:", string(char))
	}

	// Conversión a mayúsculas y minúsculas

	cadena4 := "Hola Mundo"
	fmt.Println("Mayúsculas:", strings.ToUpper(cadena4)) // Imprime: "HOLA MUNDO"
	fmt.Println("Minúsculas:", strings.ToLower(cadena4)) // Imprime: "hola mundo"

	// Reemplazo

	cadena5 := "Hola Mundo"
	nuevaCadena := strings.Replace(cadena5, "Mundo", "Gopher", 1)
	fmt.Println("Cadena con reemplazo:", nuevaCadena) // Imprime: "Hola Gopher"

	// División

	cadena6 := "Hola,Mundo,Gopher"
	partes := strings.Split(cadena6, ",")
	fmt.Println("División de cadena:", partes) // Imprime: ["Hola" "Mundo" "Gopher"]

	// Unión

	partesUnidas := strings.Join(partes, "-")
	fmt.Println("Unión de partes:", partesUnidas) // Imprime: "Hola-Mundo-Gopher"

	// Interpolación

	nombre := "Gopher"
	edad := 5
	mensaje := fmt.Sprintf("Hola, soy %s y tengo %d años.", nombre, edad)
	fmt.Println("Mensaje interpolado:", mensaje) // Imprime: "Hola, soy Gopher y tengo 5 años."

	// Verificación

	cadena7 := "Hola Mundo"
	contieneHola := strings.Contains(cadena7, "Hola")
	fmt.Println("¿La cadena contiene 'Hola'?", contieneHola) // Imprime: true

	// Comparación de cadenas

	cadena8 := "Hola"
	cadena9 := "Hola"
	fmt.Println("¿cadena1 es igual a cadena2?", cadena8 == cadena9) // Imprime: true


	// Elimina los caracteres especificados del principio y el final de una cadena

	cadena10 := "   ¡¡¡Hola Mundo!!!   "
	cadenaLimpia := strings.Trim(cadena10, " ¡")
	fmt.Println("Cadena limpia:", cadenaLimpia) // Imprime: "Hola Mundo"

	// Contar Ocurrencias

	cadena11 := "abracadabra"
	ocurrencias := strings.Count(cadena11, "a")
	fmt.Println("Número de ocurrencias de 'a':", ocurrencias) // Imprime: 5

	// Buscar subcadenas 

	cadena12 := "Hola Mundo"
	posicion := strings.Index(cadena12, "Mundo")
	fmt.Println("Posición de 'Mundo':", posicion) // Imprime: 5

	// verifican si una cadena comienza o termina con un prefijo o sufijo específico

	cadena13 := "Hola Mundo"
	fmt.Println("¿La cadena comienza con 'Hola'?", strings.HasPrefix(cadena13, "Hola"))   // Imprime: true
	fmt.Println("¿La cadena termina con 'Mundo'?", strings.HasSuffix(cadena13, "Mundo")) // Imprime: true

	// Elimina los espacios en blanco iniciales y finales de una cadena

	cadena14 := "   Hola Mundo   "
	cadenaRecortada := strings.TrimSpace(cadena14)
	fmt.Println("Cadena recortada:", cadenaRecortada) // Imprime: "Hola Mundo"

	// strings.Fields en Go se utiliza para dividir una cadena en palabras individuales

	cadena15 := "Hola   Mundo  Gopher"
  palabras := strings.Fields(cadena15)
  fmt.Println("Palabras encontradas:", palabras) // Imprime: ["Hola" "Mundo" "Gopher"]
}