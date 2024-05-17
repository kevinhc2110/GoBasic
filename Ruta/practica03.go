package main

import "fmt"

var contactos = map[string]string {"Romeo":"3179812397", "Sofia": "3207783344"}

var nombre string
var numero string

func main() {

	var option int
	fmt.Println("Ingrese una opción:")
	fmt.Scanln(&option)

	switch option {
	case 1:
		insertarContacto()	
	}

	fmt.Println(contactos)
	
}

func insertarContacto() {

	fmt.Println("Ingrese el numero")
	fmt.Scanln(&numero)

	fmt.Println("Ingrese el nombre:")
	fmt.Scanln(&nombre)

	contactos[nombre] = numero;
}

