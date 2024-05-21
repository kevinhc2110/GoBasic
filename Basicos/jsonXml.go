package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Estructura para almacenar los datos
// Las etiquetas json y xml se utilizan para especificar los nombres de los campos JSON y XML, respectivamente, al serializar y deserializar los datos.

type Datos struct {
	Nombre                 string   `json:"nombre" xml:"Nombre"`
	Edad                   int      `json:"edad" xml:"Edad"`
	FechaDeNacimiento      string   `json:"fecha_de_nacimiento" xml:"FechaDeNacimiento"`
	LenguajesDeProgramacion []string `json:"lenguajes_de_programacion" xml:"LenguajesDeProgramacion>Lenguaje"`
}

const (
	jsonFile = "datos.json"
	xmlFile  = "datos.xml"
)

// Función para crear y guardar los datos en un archivo JSON
// Abre el archivo JSON para escritura usando os.Create.
// Crea un codificador JSON usando json.NewEncoder.
// Establece la sangría para la salida JSON usando encoder.SetIndent.
// Codifica la estructura datos en el archivo JSON usando encoder.Encode.

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

// Función para borrar un archivo

func borrarArchivo(filename string) error {
	return os.Remove(filename)
}

func jsonXml()/*main*/ {

	// Crea una instancia de estructura Datos con datos de muestra.

	datos := Datos{
		Nombre:                 "Juan Perez",
		Edad:                   30,
		FechaDeNacimiento:      "1993-05-15",
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

	fmt.Println("Archivos JSON y XML creados, mostrados y borrados exitosamente.")
}
