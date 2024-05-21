package main

import (
	"fmt"
)

// DatosPersonales representa la estructura de los datos personales
type DatosPersonales struct {
	Name                string
	Age                 int
	BirthDate           string
	ProgrammingLanguages []string
}

// NewDatosPersonales crea una nueva instancia de DatosPersonales
func NewDatosPersonales(name string, age int, birthDate string, languages []string) DatosPersonales {
	return DatosPersonales{
		Name:                name,
		Age:                 age,
		BirthDate:           birthDate,
		ProgrammingLanguages: languages,
	}
}

func main() {
	// Crear una instancia de DatosPersonales
	datos := NewDatosPersonales("Tu Nombre", 30, "01/01/1990", []string{"Go", "Python", "Java"})

	// Mostrar los datos
	fmt.Printf("Datos Personales: %+v\n", datos)
}
