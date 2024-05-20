package main

import (
	"fmt"
)

// Definición de una interfaz
type Saludador interface {
    Saludar()
}

// Definición de un struct Persona
type Persona2 struct {
    Nombre string
}

// Implementación del método Saludar para Persona
func (p2 Persona2) Saludar() {
    fmt.Printf("Hola, mi nombre es %s.\n", p2.Nombre)
}

func interfaces() /*main*/ {
    // Creación de una instancia de Persona
    p2 := Persona2{
        Nombre: "Juan",
    }

    var s Saludador = p2
    s.Saludar()
}