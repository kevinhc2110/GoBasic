package main

import (
	"bytes" // Paquete que proporciona funciones para manipular secuencias de bytes
	"fmt"   // Paquete para la salida estándar, como imprimir mensajes
	"io"    // Paquete que ofrece herramientas para manejo de entrada y salida (como leer archivos)
	"os"    // Paquete para interactuar con el sistema operativo (lectura y escritura de archivos)
)

const (
	blockSize     = 512 // Tamaño del bloque que se leerá del archivo, en bytes
	jpegHeaderLen = 4   // Tamaño del encabezado de un archivo JPEG (4 bytes)
)

// Firma JPEG: los primeros 3 bytes que identifican un archivo JPEG
var jpegSignature = []byte{0xff, 0xd8, 0xff}

func main() {
	// Verificamos que el programa reciba exactamente 1 argumento (el nombre del archivo .raw)
	if len(os.Args) != 2 {
		fmt.Println("Usage: recover image.raw") // Mensaje de uso correcto
		return
	}

	// Abrimos el archivo de imagen (.raw) usando el nombre proporcionado
	imageFilename := os.Args[1]
	file, err := os.Open(imageFilename)
	if err != nil {
		fmt.Printf("Error: Could not open file %s\n", imageFilename) // Si no se puede abrir, mostramos un error
		return
	}
	defer file.Close() // Cerramos el archivo automáticamente cuando termine la función

	// Llamamos a la función para recuperar los JPEGs
	recoverJPEGs(file)
}

// recoverJPEGs se encarga de leer el archivo .raw y extraer los archivos JPEG
func recoverJPEGs(file *os.File) {
	var imgFile *os.File              // Variable para el archivo JPEG que se va a escribir
	buffer := make([]byte, blockSize) // Buffer para almacenar los bloques de 512 bytes leídos
	fileCount := 0                    // Contador de archivos JPEG recuperados (para nombrarlos secuencialmente)
	foundJPEG := false                // Flag que indica si hemos encontrado el inicio de un archivo JPEG

	// Bucle para leer el archivo de imagen en bloques de 512 bytes
	for {
		// Leemos el siguiente bloque de datos del archivo
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF { // Si hay un error distinto a llegar al final del archivo, mostramos un mensaje
			fmt.Println("Error reading from file:", err)
			return
		}

		// Si no se han leído bytes (fin del archivo), salimos del bucle
		if bytesRead == 0 {
			break
		}

		// Verificamos si el bloque contiene la firma de un archivo JPEG
		if isJPEGHeader(buffer) {
			// Si ya estábamos escribiendo en un archivo JPEG, lo cerramos
			if foundJPEG && imgFile != nil {
				imgFile.Close()
			}

			// Creamos un nuevo archivo JPEG con el nombre secuencial (000.jpg, 001.jpg, ...)
			imgFilename := fmt.Sprintf("%03d.jpg", fileCount)
			imgFile, err = os.Create(imgFilename)
			if err != nil {
				fmt.Printf("Error: Could not create file %s\n", imgFilename) // Si no podemos crear el archivo, mostramos un error
				return
			}
			fileCount++      // Incrementamos el contador para el próximo archivo
			foundJPEG = true // Marcamos que hemos encontrado un nuevo archivo JPEG
		}

		// Si hemos encontrado un archivo JPEG, escribimos el bloque leído en el archivo
		if foundJPEG && imgFile != nil {
			_, err = imgFile.Write(buffer[:bytesRead]) // Escribimos solo los bytes leídos
			if err != nil {
				fmt.Println("Error writing to JPEG file:", err) // Si hay un error al escribir, lo mostramos
				return
			}
		}
	}

	// Si un archivo JPEG estaba abierto, lo cerramos al final
	if imgFile != nil {
		imgFile.Close()
	}
}

// isJPEGHeader verifica si el bloque contiene la firma de un archivo JPEG
func isJPEGHeader(buffer []byte) bool {
	// Comparamos los primeros 3 bytes del bloque con la firma JPEG y verificamos que el cuarto byte sea válido
	return len(buffer) >= jpegHeaderLen && bytes.Equal(buffer[:3], jpegSignature) && (buffer[3]&0xf0) == 0xe0
}
