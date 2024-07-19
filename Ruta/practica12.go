//* JSON y XML

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
DIFICULTAD EXTRA:
Utilizando la lógica de creación de los archivos anteriores, crea un
programa capaz de leer y transformar en una misma clase custom de tu
lenguaje los datos almacenados en el XML y el JSON.
Borra los archivos.
*/

// Estructura para almacenar los datos

type Datos struct {
	Nombre                  string   `json:"nombre" xml:"Nombre"`
	Edad                    int      `json:"edad" xml:"Edad"`
	FechaDeNacimiento       string   `json:"fecha_de_nacimiento" xml:"FechaDeNacimiento"`
	LenguajesDeProgramacion []string `json:"lenguajes_de_programacion" xml:"LenguajesDeProgramacion>Lenguaje"`
}

const (
	jsonFile = "datos.json"
	xmlFile  = "datos.xml"
)

// Función para crear y guardar los datos en un archivo JSON

func crearArchivoJSON(datos Datos) error {
	file, err := os.Create(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(datos)
}

// Función para crear y guardar los datos en un archivo XML

func crearArchivoXML(datos Datos) error {
	file, err := os.Create(xmlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	return encoder.Encode(datos)
}

// Función para leer y mostrar el contenido de un archivo

func mostrarContenidoArchivo(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Printf("Contenido de %s:\n%s\n", filename, string(data))
	return nil
}

// Función para leer datos de un archivo JSON
// var datos Datos: Declara una variable datos de tipo Datos para almacenar los datos analizados.
// data, err := ioutil.ReadFile(filename): Lee el contenido del archivo especificado filename en la porción de bytes data. Maneja posibles errores durante la lectura del archivo.
// err := json.Unmarshal(data, &datos): Intenta deserializar los datos JSON en data en la estructura datos. Maneja posibles errores durante la deserialización.
// return datos, err: Devuelve la estructura datos que contiene los datos analizados y cualquier error encontrado durante el proceso.

func leerArchivoJSON(filename string) (Datos, error) {
	var datos Datos
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return datos, err
	}
	err = json.Unmarshal(data, &datos)
	return datos, err
}

// Función para leer datos de un archivo XML

func leerArchivoXML(filename string) (Datos, error) {
	var datos Datos
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return datos, err
	}
	err = xml.Unmarshal(data, &datos)
	return datos, err
}

// Función para borrar un archivo

func borrarArchivo(filename string) error {
	return os.Remove(filename)
}

func practica12() /*main*/ {
	datos := Datos{
		Nombre:                  "Juan Perez",
		Edad:                    30,
		FechaDeNacimiento:       "1993-05-15",
		LenguajesDeProgramacion: []string{"Go", "Python", "JavaScript"},
	}

	// Crear archivo JSON

	err := crearArchivoJSON(datos)
	if err != nil {
		fmt.Println("Error creando archivo JSON:", err)
		return
	}

	// Crear archivo XML

	err = crearArchivoXML(datos)
	if err != nil {
		fmt.Println("Error creando archivo XML:", err)
		return
	}

	// Mostrar contenido del archivo JSON

	err = mostrarContenidoArchivo(jsonFile)
	if err != nil {
		fmt.Println("Error mostrando contenido del archivo JSON:", err)
		return
	}

	// Mostrar contenido del archivo XML

	err = mostrarContenidoArchivo(xmlFile)
	if err != nil {
		fmt.Println("Error mostrando contenido del archivo XML:", err)
		return
	}

	// Leer datos del archivo JSON

	datosJSON, err := leerArchivoJSON(jsonFile)
	if err != nil {
		fmt.Println("Error leyendo archivo JSON:", err)
		return
	}
	fmt.Printf("Datos leídos del archivo JSON: %+v\n", datosJSON)

	// Leer datos del archivo XML

	datosXML, err := leerArchivoXML(xmlFile)
	if err != nil {
		fmt.Println("Error leyendo archivo XML:", err)
		return
	}
	fmt.Printf("Datos leídos del archivo XML: %+v\n", datosXML)

	// Borrar archivo JSON

	err = borrarArchivo(jsonFile)
	if err != nil {
		fmt.Println("Error borrando archivo JSON:", err)
		return
	}

	// Borrar archivo XML

	err = borrarArchivo(xmlFile)
	if err != nil {
		fmt.Println("Error borrando archivo XML:", err)
		return
	}

	fmt.Println("Archivos JSON y XML creados, leídos, mostrados y borrados exitosamente.")
}
