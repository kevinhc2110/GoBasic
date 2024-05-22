package main

// Las expresiones regulares (Regex) son patrones que se utilizan para buscar y manipular texto. Son herramientas poderosas y versátiles que se encuentran en la mayoría de los lenguajes de programación, incluido Golang.
// Patrones: Las expresiones regulares están compuestas por patrones que representan secuencias de caracteres, caracteres individuales o metacaracteres con significados especiales.
// Metacaracteres: Los metacaracteres tienen un significado especial dentro de las expresiones regulares, como . (cualquier carácter), * (cero o más repeticiones), + (una o más repeticiones), ? (cero o una repetición), ^ (comienzo de la cadena), $ (fin de la cadena), [] (rango de caracteres), etc.
// Motores de Regex: Golang utiliza motores de expresiones regulares para interpretar y ejecutar los patrones. El paquete regexp proporciona funciones para trabajar con expresiones regulares.

// Búsqueda de Patrones:
// La función regexp.MatchString(regex, text) devuelve true si el texto coincide con el patrón regular especificado por regex.
// La función regexp.FindString(regex, text) devuelve la primera subcadena que coincide con el patrón regular en el texto.
// La función regexp.FindAllString(regex, text, n) devuelve una lista de subcadenas que coinciden con el patrón regular en el texto, hasta un máximo de n coincidencias.

//Extracción de Información:
// La función regexp.Match(regex, text) devuelve un objeto *regexp.Match que contiene información sobre las coincidencias, como las subcadenas coincidentes y sus posiciones.
// Los grupos de captura permiten extraer subcadenas específicas dentro del patrón regular. Se utilizan paréntesis para definir grupos en el patrón, y luego se accede a ellos mediante índices o nombres dentro de la expresión.

// Manipulación de Texto:
// La función regexp.ReplaceAllString(regex, replacement, text) reemplaza todas las coincidencias del patrón regular en el texto por la cadena de reemplazo especificada.
// La función regexp.Func.ReplaceString(regex, func(match *regexp.Match) string, text) permite realizar reemplazos personalizados utilizando una función de reemplazo.

import (
	"fmt"
	"regexp"
)

func expresionesRegulares() {

	// Patrón para extraer números de teléfono
	phoneRegex := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)

	// Texto de ejemplo
	text := "Mi número de teléfono es 123-456-7890. También puedes contactarme al 987-654-3210."

	// Búsqueda de coincidencias
	matches := phoneRegex.FindAllString(text, -1)
	fmt.Println("Números de teléfono encontrados:", matches)

	// Extracción de grupos de captura
	for _, match := range matches {
			fmt.Printf("Número: %s\n", match[1:]) // Extracción de grupos de captura (digitos)
	}

	// Reemplazo de texto
	newText := phoneRegex.ReplaceAllString(text, "[REDACTED]")
	fmt.Println("Texto con números reemplazados:", newText)

	// Implementación función encontrarNumeros

	texto := "El precio es $25.50 y hay 100 unidades disponibles."
	numeros := encontrarNumeros(texto)
	fmt.Println("Números encontrados:", numeros)

	// Mas ejemplos

	// MatchString determina si una cadena coincide con el patrón de la expresión regular
	// FindAllString encuentra todas las subcadenas que coinciden con el patrón
	// ReplaceAllString reemplaza todas las coincidencias de un patrón por otra cadena

}

func encontrarNumeros(texto string) []string {

	// Definir la expresión regular para encontrar números pero con panic, también podemos usar compile

	expresion := regexp.MustCompile(`\d+`)

	// Encontrar todos los números en el texto

	numeros := expresion.FindAllString(texto, -1)
	return numeros
}

// Patrones importantes
// \d: Coincide con un dígito (equivalente a [0-9]).
// \D: Coincide con cualquier carácter que no sea un dígito (equivalente a [^0-9]).
// \w: Coincide con un carácter de palabra (letras, dígitos o guiones bajos).
// \W: Coincide con cualquier carácter que no sea un carácter de palabra.
// \s: Coincide con un espacio en blanco (espacio, tabulación, salto de línea).
// \S: Coincide con cualquier carácter que no sea un espacio en blanco.
// .: Coincide con cualquier carácter excepto un salto de línea.
// ^: Coincide con el inicio de una cadena.
// $: Coincide con el final de una cadena.
// [abc]: Coincide con cualquier carácter entre corchetes (a, b o c).
// [^abc]: Coincide con cualquier carácter que no esté entre corchetes.
// [a-z]: Coincide con cualquier carácter entre a y z en minúsculas.
// [A-Z]: Coincide con cualquier carácter entre A y Z en mayúsculas.
// [0-9]: Coincide con cualquier dígito.
// \b: Coincide con un límite de palabra (inicio o fin de una palabra).

// Ejemplo de combinaciones
// \d{3}: Coincide con tres dígitos consecutivos.
// \w+: Coincide con una o más letras, dígitos o guiones bajos.
// ^\d{3}-\d{3}-\d{4}$: Coincide con un número de teléfono en formato ###-###-####.
// \b[A-Z]\w\b:* Coincide con palabras que comienzan con una letra mayúscula.
// \d{2}/\d{2}/\d{4}: Coincide con fechas en formato dd/mm/yyyy.
// ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$: Validación de formato de correo electrónico
// (https?|ftp)://[^\s/$.?#].[^\s]*: Extracción de URL de un texto
// \b(palabra)\b: Búsqueda de palabras específicas
// ^(\+\d{1,2}\s?)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$: Validación de números de teléfono
// (?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+}{"':;?/>.<,])(?=.*[a-zA-Z]).{8,}: Validación de contraseñas 
// \b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b: Extracción de direcciones IP
// ^a\w+: Buscar todas las palabras que comiencen con "a" en un texto
// [\w\.-]+@[\w\.-]+\.[a-z]{2,6}: Extraer las direcciones de correo electrónico de un documento
// sed -i 's/hola/adiós/g' archivo.txt: Reemplazar todas las ocurrencias de "hola" por "adiós" en un archivo


// *También podemos usar clases POSIX*
// [:alnum:]	alfanumérico (≡ [0-9A-Za-z])
// [:alpha:]	alfabético (≡ [A-Za-z])
// [:ascii:]	ASCII (≡ [\x00-\x7F])
// [:blank:]	en blanco (≡ [\t ])
// [:cntrl:]	control (≡ [\x00-\x1F\x7F])
// [:digit:]	dígitos (≡ [0-9])
// [:graph:]	gráfico (≡ [!-~] == [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~])
// [:lower:]	minúsculas (≡ [a-z])
// [:print:]	imprimible (≡ [ -~] == [ [:graph:]])
// [:punct:]	puntuación (≡ [!-/:-@[-`{-~])
// [:space:]	espacio en blanco (≡ [\t\n\v\f\r ])
// [:upper:]	mayúsculas (≡ [A-Z])
// [:word:]	caracteres de palabras (≡ [0-9A-Za-z_])
// [:xdigit:]	dígito hexadecial (≡ [0-9A-Fa-f])



