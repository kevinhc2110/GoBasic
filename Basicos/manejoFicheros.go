package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func manejoFicheros() {

	// Obtener el nombre de usuario de GitHub

	githubUsername := "kevinhc2110"

	// Crear el nombre de archivo con el nombre de usuario de GitHub y extensión .txt

	filename := githubUsername + ".txt"

	// Contenido a escribir en el archivo
	// fmt.Sprintf es una función del paquete fmt que formatea una cadena de texto según un formato específico y devuelve la cadena resultante. A diferencia de fmt.Printf, que imprime directamente en la salida estándar, fmt.Sprintf devuelve la cadena formateada.

	content := fmt.Sprintf(`Nombre: %s
	Edad: %d
	Lenguaje de programación favorito: %s
	`, "Kevin", 25, "Go")

	// Escribir contenido en el archivo
	// ioutil.WriteFile es una función del paquete io/ioutil que escribe datos en un archivo. Si el archivo no existe, WriteFile lo crea con los permisos especificados. Si el archivo ya existe, su contenido se sobrescribe.

	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	fmt.Println("Archivo creado exitosamente.")

	// Leer y mostrar el contenido del archivo

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	fmt.Println("Contenido del archivo:")
	fmt.Println(string(data))

	// Borrar el archivo
	// os.Remove es una función del paquete os que elimina (borra) un archivo o un directorio vacío.

	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error al borrar el archivo:", err)
		return
	}
	fmt.Println("Archivo borrado exitosamente.")
}