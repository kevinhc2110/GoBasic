package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack []string

type Queue []string

func main() {

	//En la función main, se crean dos pilas, backStack y forwardStack, para rastrear las páginas visitadas hacia atrás y hacia adelante respectivamente

	backStack := Stack{}     // Pila para las páginas visitadas hacia atrás
	forwardStack := Stack{}  // Pila para las páginas visitadas hacia adelante
	currentPage := "Inicio" // Página actual

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenido al navegador web simple.")

	for {
		fmt.Println("Página actual:", currentPage)

		fmt.Print("Ingrese una acción (adelante, atrás o el nombre de una página): ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Eliminar el salto de línea al final

		switch input {
		case "adelante":
			if page := forwardStack.Pop(); page != "" {
				backStack.Push(currentPage)
				currentPage = page
				fmt.Println("Página adelante:", currentPage)
			} else {
				fmt.Println("No hay páginas adelante.")
			}
		case "atrás":
			if page := backStack.Pop(); page != "" {
				forwardStack.Push(currentPage)
				currentPage = page
				fmt.Println("Página atrás:", currentPage)
			} else {
				fmt.Println("No hay páginas atrás.")
			}
		case "salir":
			fmt.Println("Saliendo del navegador...")
			break
		default:
			backStack.Push(currentPage) // Agregar la página actual a la pila de atrás
			currentPage = input         // Establecer la página ingresada como la nueva página actual
			forwardStack = Stack{}      // Vaciar la pila de adelante
			fmt.Println("Navegando a la página:", input)
		}
	}
}

// func (s *Stack) Push(page string) {: Esto define una función llamada Push que pertenece al tipo Stack. El receptor de este método es un puntero a Stack, denotado por (s *Stack), lo que significa que esta función se puede llamar en cualquier instancia de Stack y puede modificar el valor de la instancia original. El parámetro page string indica que la función toma un argumento de tipo string llamado page.

func (s *Stack) Push(page string) { 
	*s = append(*s, page)
}

// if len(*s) == 0 { return "" }: Esta línea verifica si la longitud del slice *s (el slice de strings al que se apunta) es cero, lo que significa que la pila está vacía. Si la pila está vacía, la función devuelve una cadena vacía "" para indicar que no hay nada que sacar de la pila.
// top := (*s)[len(*s)-1]: Si la pila no está vacía, esta línea obtiene el elemento superior de la pila, que es el último elemento del slice *s. El elemento superior se almacena en la variable top.
//: *s = (*s)[:len(*s)-1]: Luego, esta línea actualiza la pila eliminando el elemento superior. Esto se hace creando un nuevo slice que contiene todos los elementos del slice original *s, excepto el último elemento. El puntero s se actualiza para que apunte al nuevo slice, lo que efectivamente elimina el elemento superior de la pila.
//return top: Finalmente, la función devuelve el elemento superior que se extrajo de la pila.

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
