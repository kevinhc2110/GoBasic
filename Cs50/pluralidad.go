package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//* Problema 2 Plurality

// Definición de constantes y estructuras

// Número máximo de candidatos
const MAX = 9

// Estructura para representar un candidato
type Candidato struct {
	Nombre string // Nombre del candidato
	Votos  int    // Cantidad de votos que ha recibido
}

// Arreglo de candidatos
var Candidatos [MAX]Candidato

// Número de candidatos
var NumeroDeCandidatos int

func pluralidad() {
	// Verificación de uso válido del programa
	if len(os.Args) < 2 {
		fmt.Println("Uso: plurality [candidato ...]")
		os.Exit(1)
	}

	// Poblar el arreglo de candidatos desde los argumentos de la línea de comandos
	// Se le debe restar 1 porque os.Args incluye el nombre del programa como su primer elemento, seguido por los argumentos de la línea de comandos.
	NumeroDeCandidatos = len(os.Args) - 1
	if NumeroDeCandidatos > MAX {
		fmt.Printf("El número máximo de candidatos es %d\n", MAX)
		os.Exit(2)
	}
	for i := 0; i < NumeroDeCandidatos; i++ {
		// Se debe sumar 1 por que el primer valor del array de argumentos es el nombre del programa
		Candidatos[i].Nombre = os.Args[i+1]
		Candidatos[i].Votos = 0
	}

	// Obtener el número de votantes desde la entrada estándar
	numeroDeVotantes := ObtenerEntero("Número de votantes: ")

	// Iterar sobre todos los votantes
	for i := 0; i < numeroDeVotantes; i++ {
		nombre := ObtenerCadena("Voto: ")

		// Verificar si el voto es válido
		if !votar(nombre) {
			fmt.Println("Voto inválido.")
		}
	}

	// Mostrar el ganador (o ganadores) de la elección
	imprimirGanador()
}

// Funciones auxiliares

// Función para leer un entero desde la entrada estándar
func ObtenerEntero(prompt string) int {
	var num int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&num)
		if err == nil {
			return num
		}
		fmt.Println("Entrada inválida. Por favor, ingrese un número entero.")
		// Limpiar el buffer de entrada
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

// Función para leer una cadena desde la entrada estándar
func ObtenerCadena(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Función para actualizar los totales de votos dado un nuevo voto
func votar(nombre string) bool {
	for i := 0; i < NumeroDeCandidatos; i++ {
		if Candidatos[i].Nombre == nombre {
			Candidatos[i].Votos++
			return true
		}
	}
	return false // Devuelve false si el nombre no coincide con ningún candidato registrado
}

// Función para imprimir el ganador (o ganadores) de la elección
func imprimirGanador() {
	// Encontrar el máximo número de votos obtenidos
	maxVotos := Candidatos[0].Votos
	for i := 1; i < NumeroDeCandidatos; i++ {
		if Candidatos[i].Votos > maxVotos {
			maxVotos = Candidatos[i].Votos
		}
	}

	// Imprimir todos los candidatos que tienen el máximo número de votos
	for i := 0; i < NumeroDeCandidatos; i++ {
		if Candidatos[i].Votos == maxVotos {
			fmt.Println(Candidatos[i].Nombre)
		}
	}
}
