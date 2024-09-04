package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// Constante para el tamaño del encabezado del archivo WAV
const headerSize = 44

func main() {
	// Verificar argumentos de la línea de comandos
	if len(os.Args) != 4 {
		log.Fatalf("Uso: %s input.wav output.wav factor\n", os.Args[0])
	}

	// Abrir el archivo de entrada
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo de entrada: %v", err)
	}
	defer inputFile.Close()

	// Crear el archivo de salida
	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("No se pudo crear el archivo de salida: %v", err)
	}
	defer outputFile.Close()

	// Obtener el factor de escala
	factor, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("No se pudo interpretar el factor: %v", err)
	}

	// Leer y escribir el encabezado (44 bytes)
	header := make([]byte, headerSize)
	_, err = inputFile.Read(header)
	if err != nil {
		log.Fatalf("No se pudo leer el encabezado: %v", err)
	}

	_, err = outputFile.Write(header)
	if err != nil {
		log.Fatalf("No se pudo escribir el encabezado: %v", err)
	}

	// Leer las muestras, escalarlas y escribirlas en el archivo de salida
	var sample int16
	for {
		// Leer una muestra (16 bits)
		err = binary.Read(inputFile, binary.LittleEndian, &sample)
		if err == io.EOF {
			break // Llegamos al final del archivo
		}
		if err != nil {
			log.Fatalf("Error leyendo muestra: %v", err)
		}

		// Escalar la muestra
		scaledSample := float64(sample) * factor

		// Limitar el valor a 16 bits con signo
		if scaledSample > math.MaxInt16 {
			scaledSample = math.MaxInt16
		} else if scaledSample < math.MinInt16 {
			scaledSample = math.MinInt16
		}

		// Convertir de nuevo a int16
		sample = int16(scaledSample)

		// Escribir la muestra escalada en el archivo de salida
		err = binary.Write(outputFile, binary.LittleEndian, sample)
		if err != nil {
			log.Fatalf("Error escribiendo muestra: %v", err)
		}
	}

	fmt.Println("El archivo de salida ha sido generado correctamente.")
}
