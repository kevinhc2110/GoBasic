package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Lee el fichero "Challenge21.txt" incluido en el proyecto, calcula su
resultado e imprímelo.
- El .txt se corresponde con las entradas de una calculadora.
- Cada línea tendrá un número o una operación representada por un
símbolo (alternando ambos).
- Soporta números enteros y decimales.
- Soporta las operaciones suma "+", resta "-", multiplicación "
y división "/".
- El resultado se muestra al finalizar la lectura de la última
línea (si el .txt es correcto).
- Si el formato del .txt no es correcto, se indicará que no se han
podido resolver las operaciones.
*/

// Tipos de error personalizados para errores específicos
type ErrorFormatoInvalido struct {
	mensaje string
}

type ErrorDivisionPorCero struct{}

func (err ErrorFormatoInvalido) Error() string {
	return err.mensaje
}

func (err ErrorDivisionPorCero) Error() string {
	return "Error: División por cero"
}

func calculadoraTXT() {
	// Abre el archivo
	archivo, err := os.Open("calculadoraTXT.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	// Inicializa variables
	var resultado float64
	var operador string
	primeraLinea := true
	lector := bufio.NewScanner(archivo)

	// Lee el archivo línea por línea
	for lector.Scan() {
		linea := strings.TrimSpace(lector.Text())

		if linea == "" {
			continue
		}

		if primeraLinea {
			// La primera línea debe ser un número
			resultado, err = strconv.ParseFloat(linea, 64)
			if err != nil {
				fmt.Println("Formato de archivo inválido: la primera línea no es un número")
				return
			}
			primeraLinea = false
		} else {
			if operador == "" {
				// La siguiente línea debe ser un operador
				if linea == "+" || linea == "-" || linea == "*" || linea == "/" {
					operador = linea
				} else {
					err := ErrorFormatoInvalido{mensaje: "Formato de archivo inválido: operador no válido"}
					fmt.Println(err)
					return
				}
			} else {
				// La siguiente línea debe ser un número
				numero, err := strconv.ParseFloat(linea, 64)
				if err != nil {
					fmt.Println("Formato de archivo inválido: se esperaba un número")
					return
				}

				// Realiza la operación
				switch operador {
				case "+":
					resultado += numero
				case "-":
					resultado -= numero
				case "*":
					resultado *= numero
				case "/":
					if numero == 0 {
						err := ErrorDivisionPorCero{}
						fmt.Println(err)
						return
					}
					resultado /= numero
				}

				// Reinicia el operador para la siguiente iteración
				operador = ""
			}
		}
	}

	// Verifica si hubo algún error al escanear el archivo
	if err := lector.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Verifica si el archivo terminó con un operador sin un número
	if operador != "" {
		err := ErrorFormatoInvalido{mensaje: "Formato de archivo inválido: terminó con un operador"}
		fmt.Println(err)
		return
	}

	// Imprime el resultado final
	fmt.Printf("El resultado es: %.2f\n", resultado)
}
