package main

import (
	"context"
	"fmt"
	"time"
)

func contexto() {

	// Crear un contexto con cancelación

	ctx, cancel := context.WithCancel(context.Background())

	// Lanzar una goroutine para realizar una tarea que puede tardar mucho tiempo

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Tarea cancelada.")
				return
			default:
				// Simular una tarea que lleva tiempo
				time.Sleep(1 * time.Second)
				fmt.Println("Realizando tarea...")
			}
		}
	}()

	// Simular una acción del usuario que desencadena la cancelación

	time.Sleep(3 * time.Second)
	fmt.Println("Cancelando tarea...")
	cancel() // Cancelar la tarea

	// Esperar un momento para que la goroutine tenga tiempo de finalizar

	time.Sleep(1 * time.Second)
	fmt.Println("Programa terminado.")
}

// En Go, el paquete context proporciona un mecanismo para la gestión de la cancelación de operaciones, el paso de valores entre llamadas de función y la gestión de deadlocks en aplicaciones concurrentes y paralelas.
// Un contexto (context.Context) representa el entorno de una solicitud, transacción o trabajo en curso. Proporciona métodos para cancelar operaciones, temporizar operaciones y almacenar y pasar valores a través de la jerarquía de llamadas de función de manera segura
// Un contexto se puede crear utilizando el método context.Background() como un contexto de fondo sin valores asociados, o utilizando context.WithCancel(), context.WithDeadline(), context.WithTimeout() o context.WithValue() para crear un contexto con funcionalidades adicionales.

// *context.TODO()

// context.TODO() devuelve un contexto sin valor y sin cancelación. Se utiliza cuando se necesita un contexto pero aún no se ha determinado su tipo específico o no se necesita ningún valor asociado con él.
// Este contexto se puede utilizar como marcador de posición cuando se espera que se proporcione un contexto específico más adelante en la función.

/*package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()
	fmt.Println(ctx)
}*/

// *context.Background()

// context.Background() devuelve un contexto sin valor y sin cancelación, similar a context.TODO(). Sin embargo, se utiliza más comúnmente en situaciones donde se necesita un contexto pero no se requiere ningún valor asociado con él y no se espera que se cancele.
// Este contexto se puede utilizar como el contexto raíz para aplicaciones o para operaciones de bajo nivel donde no se necesita control de cancelación.

/*package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
}*/