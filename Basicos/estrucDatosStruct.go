package main

import "fmt"

// Definimos una estructura llamada Persona
type Persona4 struct {
	Nombre  string
	Edad    int
	Correo  string
}

func estruClase() /*main*/ {

	// No hay clases si no estructuras

	// Crear una nueva instancia de Persona

	persona := NuevaPersona("Juan Pérez", 30, "juan.perez@example.com")

	// Imprimir los detalles de la persona

	fmt.Println("Detalles iniciales de la persona:")
	persona.ImprimirDetalles()

	// Modificar los detalles de la persona

	persona.ModificarDetalles("María López", 25, "maria.lopez@example.com")

	// Imprimir los detalles modificados de la persona
	
	fmt.Println("\nDetalles modificados de la persona:")
	persona.ImprimirDetalles()
	
}

// Método para inicializar una nueva Persona

func NuevaPersona(nombre string, edad int, correo string) *Persona4 {
	return &Persona4{
		Nombre: nombre,
		Edad:   edad,
		Correo: correo,
	}
}

// Método para imprimir los detalles de la Persona

func (p *Persona4) ImprimirDetalles() {
	fmt.Printf("Nombre: %s\nEdad: %d\nCorreo: %s\n", p.Nombre, p.Edad, p.Correo)
}

// Método para modificar los detalles de la Persona

func (p *Persona4) ModificarDetalles(nombre string, edad int, correo string) {
	p.Nombre = nombre
	p.Edad = edad
	p.Correo = correo
}


