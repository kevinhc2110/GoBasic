//* Pilas y colas

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
DIFICULTAD EXTRA:
- Utilizando la implementación de pila y cadenas de texto, simula el mecanismo adelante/atrás
de un navegador web. Crea un programa en el que puedas navegar a una página o indicarle
que te quieres desplazar adelante o atrás, mostrando en cada caso el nombre de la web.
Las palabras "adelante", "atrás" desencadenan esta acción, el resto se interpreta como
el nombre de una nueva web.
- Utilizando la implementación de cola y cadenas de texto, simula el mecanismo de una
impresora compartida que recibe documentos y los imprime cuando así se le indica.
La palabra "imprimir" imprime un elemento de la cola, el resto de palabras se
interpretan como nombres de documentos.
*/

// Estructura para representar una página web
type Pagina struct {
	Nombre string
	Contenido string
}

// Pila para almacenar páginas web

type Stack []Pagina


func stackPractica() /*main*/ {

	   // Inicialización de pilas

		 backStack := Stack{}
		 forwardStack := Stack{}
 
		 // Lector de entrada

		 reader := bufio.NewReader(os.Stdin)
 
		 // Mensaje de bienvenida

		 fmt.Println("Bienvenido al navegador web simple.")
 
		 for {

				 // Solicitud de acción al usuario

				 fmt.Print("Ingrese una acción (adelante, atrás o el nombre de una página): ")
				 input, _ := reader.ReadString('\n')
				 input = strings.TrimSpace(input)
 
				 // Instrucción switch para determinar la acción

				 switch input {
				 case "adelante":
						 if pagina := forwardStack.Pop(); pagina.Nombre != "" {
								 backStack.Push(pagina)
								 fmt.Println("Página adelante:", pagina.Nombre)
								 fmt.Println(pagina.Contenido) // Mostrar contenido de la página
						 } else {
								 fmt.Println("No hay páginas adelante.")
						 }
 
				 case "atras":
						 if pagina := backStack.Pop(); pagina.Nombre != "" {
								 forwardStack.Push(pagina)
								 fmt.Println("Página atrás:", pagina.Nombre)
								 fmt.Println(pagina.Contenido) // Mostrar contenido de la página
						 } else {
								 fmt.Println("No hay páginas atrás.")
						 }
 
				 default:

						 // Navegar a una nueva página

						 pagina := Pagina{Nombre: input, Contenido: cargarPagina(input)}
						 backStack.Push(pagina)
						 forwardStack = Stack{}
						 fmt.Println("Navegando a la página:", pagina.Nombre)
						 fmt.Println(pagina.Contenido) // Mostrar contenido de la página
				 }
		 }
}

// func (stack *Stack) Push(page Pagina) {: Esto define una función llamada Push que pertenece al tipo Stack. El receptor de este método es un puntero a Stack, denotado por (stack *Stack), lo que significa que esta función se puede llamar en cualquier instancia de Stack y puede modificar el valor de la instancia original. El parámetro page Pagina indica que la función toma un argumento de tipo Pagina llamado page.

func (stack *Stack) Push(page Pagina) { 
	*stack = append(*stack, page)
}

// if len(*s) == 0 { return "" }: Esta línea verifica si la longitud del slice *stack (el slice de strings al que se apunta) es cero, lo que significa que la pila está vacía. Si la pila está vacía, la función devuelve una estructura tipo pagina para indicar que no hay nada que sacar de la pila.
// ultimaPagina := (*stack)[len(*s)-1]: Si la pila no está vacía, esta línea obtiene el elemento superior de la pila, que es el último elemento del slice *stack. El elemento superior se almacena en la variable top.
//: *stack = (*s)[:len(*s)-1]: Luego, esta línea actualiza la pila eliminando el elemento superior. Esto se hace creando un nuevo slice que contiene todos los elementos del slice original *stack, excepto el último elemento. El puntero s se actualiza para que apunte al nuevo slice, lo que efectivamente elimina el elemento superior de la pila.
//return ultimaPagina: Finalmente, la función devuelve el elemento superior que se extrajo de la pila.

func (stack *Stack) Pop() Pagina {
	if len(*stack) == 0 {
		return Pagina{}
	}
	ultimaPagina := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ultimaPagina
}

// Función para cargar el contenido de una página web (simulación)

func cargarPagina(nombrePagina string) string {
	contenido, _ := os.ReadFile(nombrePagina + ".txt")
	return string(contenido)
}

//---------------------------------------------------//-------------------------------------------------------//

// Estructura para representar un documento
type Documento struct {
	Nombre string
	Contenido string
}

// Cola para almacenar documentos

type Queue []Documento

// Función principal
func queuePractica() /*main*/ {

	// Inicialización de la cola de documentos

	colaDocumentos := Queue{}

	// Lector de entrada

	reader := bufio.NewReader(os.Stdin)

	// Mensaje de bienvenida

	fmt.Println("Bienvenido al simulador de impresora compartida.")

	for {

			// Solicitud de acción al usuario

			fmt.Print("Ingrese una acción (imprimir o nombre de documento): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Instrucción switch para determinar la acción

			switch input {
			case "imprimir":
					if documento := colaDocumentos.Pop(); documento.Nombre != "" {
							imprimirDocumento(documento)
					} else {
							fmt.Println("No hay documentos para imprimir.")
					}

			default:

					// Agregar documento a la cola

					contenido, _ := os.ReadFile(input + ".txt")
					documento := Documento{Nombre: input, Contenido: string(contenido)}
					colaDocumentos.Push(documento)
					fmt.Println("Documento agregado a la cola:", documento.Nombre)
			}
	}
}


// Función para agregar un documento a la cola

func (queue *Queue) Push(documento Documento) {
	*queue = append(*queue, documento)
}

// Función para extraer el primer documento de la cola

func (queue *Queue) Pop() Documento {
	if len(*queue) == 0 {
			return Documento{}
	}

	documento := (*queue)[0]
	*queue = (*queue)[1:]
	return documento
}

// Función para imprimir un documento
func imprimirDocumento(documento Documento) {
	fmt.Println("Imprimiendo documento:", documento.Nombre)
	fmt.Println(documento.Contenido)
	fmt.Println("Impresión finalizada.")
}

