//* Enumeraciones

package main

import (
	"fmt"
)

/*
DIFICULTAD EXTRA:
Crea un pequeño sistema de gestión del estado de pedidos.
Implementa una clase que defina un pedido con las siguientes características:
- El pedido tiene un identificador y un estado.
- El estado es un Enum con estos valores: PENDIENTE, ENVIADO, ENTREGADO y CANCELADO.
- Implementa las funciones que sirvan para modificar el estado:
- Pedido enviado
- Pedido cancelado
- Pedido entregado
(Establece una lógica, por ejemplo, no se puede entregar si no se ha enviado, etc...)
- Implementa una función para mostrar un texto descriptivo según el estado actual.
- Crea diferentes pedidos y muestra cómo se interactúa con ellos.
*/

// Definimos el tipo `OrderState` como un alias para `int`.

type OrderState int

// Enumeramos los estados de los pedidos utilizando iota.

const (
	PENDIENTE OrderState = iota
	ENVIADO
	ENTREGADO
	CANCELADO
)

// Pedido representa un pedido con un identificador y un estado.

type Pedido struct {
	ID    int
	Estado OrderState
}

// Enviar marca el pedido como enviado.

func (p *Pedido) Enviar() {

	// Esta función cambia el estado del pedido a ENVIADO si el pedido está actualmente en el estado PENDIENTE.
	// Primero, verifica si el estado actual del pedido (p.Estado) es PENDIENTE.
  //Si el estado es PENDIENTE, cambia el estado a ENVIADO y muestra un mensaje indicando que el pedido ha sido enviado.

	if p.Estado == PENDIENTE {
		p.Estado = ENVIADO
		fmt.Printf("Pedido %d ha sido enviado.\n", p.ID)
	} else {
		fmt.Printf("Pedido %d no puede ser enviado en su estado actual.\n", p.ID)
	}
}

// Cancelar marca el pedido como cancelado.

func (p *Pedido) Cancelar() {

	// Esta función cambia el estado del pedido a CANCELADO si el pedido está actualmente en los estados PENDIENTE o ENVIADO.

	if p.Estado == PENDIENTE || p.Estado == ENVIADO {
		p.Estado = CANCELADO
		fmt.Printf("Pedido %d ha sido cancelado.\n", p.ID)
	} else {
		fmt.Printf("Pedido %d no puede ser cancelado en su estado actual.\n", p.ID)
	}
}

// Entregar marca el pedido como entregado.

func (p *Pedido) Entregar() {

	// Esta función cambia el estado del pedido a ENTREGADO si el pedido está actualmente en el estado ENVIADO.

	if p.Estado == ENVIADO {
		p.Estado = ENTREGADO
		fmt.Printf("Pedido %d ha sido entregado.\n", p.ID)
	} else {
		fmt.Printf("Pedido %d no puede ser entregado en su estado actual.\n", p.ID)
	}
}

// EstadoTexto devuelve una descripción del estado actual del pedido.

func (p *Pedido) EstadoTexto() string {
	switch p.Estado {
	case PENDIENTE:
		return "Pendiente"
	case ENVIADO:
		return "Enviado"
	case ENTREGADO:
		return "Entregado"
	case CANCELADO:
		return "Cancelado"
	default:
		return "Estado desconocido"
	}
}

func practica19() /*main*/ {

	// Crear algunos pedidos.

	pedido1 := Pedido{ID: 1, Estado: PENDIENTE}
	pedido2 := Pedido{ID: 2, Estado: PENDIENTE}

	// Mostrar estado inicial.

	fmt.Printf("Pedido %d está %s.\n", pedido1.ID, pedido1.EstadoTexto())
	fmt.Printf("Pedido %d está %s.\n", pedido2.ID, pedido2.EstadoTexto())

	// Intentar entregar un pedido que aún no se ha enviado.

	pedido1.Entregar()

	// Enviar el primer pedido.

	pedido1.Enviar()

	// Intentar cancelar un pedido enviado.

	pedido1.Cancelar()

	// Entregar el primer pedido.

	pedido1.Entregar()

	// Cancelar el segundo pedido.

	pedido2.Cancelar()

	// Mostrar estados finales.

	fmt.Printf("Pedido %d está %s.\n", pedido1.ID, pedido1.EstadoTexto())
	fmt.Printf("Pedido %d está %s.\n", pedido2.ID, pedido2.EstadoTexto())
}
