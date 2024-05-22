package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
DIFICULTAD EXTRA:
Crea una agenda de contactos por terminal.
- Debes implementar funcionalidades de búsqueda, inserción, actualización
y eliminación de contactos.
- Cada contacto debe tener un nombre y un número de teléfono.
- El programa solicita en primer lugar cuál es la operación que se quiere realizar,
y a continuación los datos necesarios para llevarla a cabo.
- El programa no puede dejar introducir números de teléfono no numéricos y con más
de 11 dígitos (o el número de dígitos que quieras).
- También se debe proponer una operación de finalización del programa.
*/

var contactos = map[string]string {"romeo":"3179812397", "sofia": "3207783344"}

var numero string

func practica03() /*main*/ {

	var option int

	fmt.Println("\nAgenda de Contactos")
	fmt.Println("1. Buscar contacto")
	fmt.Println("2. Agregar contacto")
	fmt.Println("3. Actualizar contacto")
	fmt.Println("4. Eliminar contacto")
	fmt.Println("5. Mostrar todos los contactos")
	fmt.Println("6. Salir")

	fmt.Scanln(&option)

	switch option {
		case 1:
			buscarContacto()
		case 2:
			agregarContacto()
		case 3:
			actualizarContacto()
		case 4:
			eliminarContacto()
		case 5:
			mostrarContactos()	
		case 6:
			fmt.Println("Saliendo del programa...")

			// Detiene ejecución

			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
	}
}

func buscarContacto()  {

	fmt.Print("Ingrese el nombre del contacto a buscar:  ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()

	nombre = strings.ToLower(nombre)

	numero, filtro := contactos[nombre]

	//%s funciona remplazando por la variable siguiente en orden
	// /n salto de linea

	if filtro {
		fmt.Printf("Número de teléfono de %s: %s\n", nombre, numero)
	} else {
		fmt.Println("El contacto no fue encontrado.")
	}

}

func agregarContacto() {

	//bufio.NewScanner(os.Stdin) y scanner.Scan(): Permite leer lineas de texto completas

	fmt.Print("Ingrese el nombre del nuevo contacto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()

	// Convertir en minúsculas una cadena de texto

	nombre = strings.ToLower(nombre)

	//fmt.Scanln: Lee la entrada hasta el primer espacio en planco

	fmt.Println("Ingrese el numero")
	fmt.Scanln(&numero)

	if !esNumero(numero) || len(numero) > 11 {
		fmt.Println("El número de teléfono ingresado no es válido.")
		return
	}

	contactos[nombre] = numero;

	fmt.Println("Contacto agregado correctamente.")
}

// strconv.Atoi(): Esto se usa para convertir una cadena de texto a string, devuelve 2 valores 
// Usamos "_" para ignorar el primer valor que es el entero
// El segundo valor es error, si se completa la conversion error == nil, si no genera un error

func esNumero(cadena string) bool {
	_, err := strconv.Atoi(cadena)
	return err == nil
}

func actualizarContacto() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese el nombre del contacto a actualizar: ")
	scanner.Scan()
	nombre := scanner.Text()

	nombre = strings.ToLower(nombre)

	// Si el contacto no fue encontrado se detiene la ejecución 
	
	_, encontrado := contactos[nombre]
	if !encontrado {
		fmt.Println("El contacto no fue encontrado.")
		return
	}

	var option int

	fmt.Println("\nIngrese un opción")
	fmt.Println("1. Actualizar numero")
	fmt.Println("2. Actualizar nombre")
	fmt.Println("3. Actualizar numero y nombre")

	fmt.Scanln(&option)

	switch option {
		case 1:
			fmt.Print("Ingrese el nuevo número de teléfono del contacto: ")
			scanner.Scan()
			nuevoNumero := scanner.Text()
		
			if !esNumero(nuevoNumero) || len(nuevoNumero) > 11 {
				fmt.Println("El número de teléfono ingresado no es válido.")
				return
			}
		
			contactos[nombre] = nuevoNumero
			fmt.Println("Contacto actualizado correctamente.")

		case 2:
			fmt.Print("Ingrese el nuevo nombre del contacto: ")
      scanner.Scan()
      nuevoNombre := strings.ToLower(scanner.Text())

      contactos[nuevoNombre] = contactos[nombre]
      delete(contactos, nombre)
      fmt.Println("Contacto actualizado correctamente.")
		
		case 3:
			fmt.Print("Ingrese el nuevo nombre del contacto: ")
			scanner.Scan()
			nuevoNombre := strings.ToLower(scanner.Text())

			fmt.Print("Ingrese el nuevo número de teléfono del contacto: ")
			scanner.Scan()
			nuevoNumero := scanner.Text()

			if !esNumero(nuevoNumero) || len(nuevoNumero) > 11 {
					fmt.Println("El número de teléfono ingresado no es válido.")
					return
			}

			contactos[nuevoNombre] = nuevoNumero
			delete(contactos, nombre)
			fmt.Println("Contacto actualizado correctamente.")

		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
	}
}

func eliminarContacto() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese el nombre del contacto a eliminar: ")
	scanner.Scan()
	nombre := scanner.Text()

	nombre = strings.ToLower(nombre)

	_, encontrado := contactos[nombre]
	if !encontrado {
		fmt.Println("El contacto no fue encontrado.")
		return
	}

	delete(contactos, nombre)
	fmt.Println("Contacto eliminado correctamente.")
}

func mostrarContactos() {
	fmt.Println("\nLista de Contactos:")
	for nombre, numero := range contactos {
		fmt.Printf("%s: %s\n", nombre, numero)
	}
}
