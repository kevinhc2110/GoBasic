package main

import (
	"errors"
	"fmt"
)

/*
DIFICULTAD EXTRA:
Crea una función que sea capaz de procesar parámetros, pero que también
pueda lanzar 3 tipos diferentes de excepciones (una de ellas tiene que
corresponderse con un tipo de excepción creada por nosotros de manera
personalizada, y debe ser lanzada de manera manual) en caso de error.
- Captura todas las excepciones desde el lugar donde llamas a la función.
- Imprime el tipo de error.
- Imprime si no se ha producido ningún error.
- Imprime que la ejecución ha finalizado.
*/

// ErrorPersonalizado es un tipo de error personalizado

type ErrorPersonalizado struct {
	message string
}

// Implementa el método Error() de la interfaz error

func (e ErrorPersonalizado) Error() string {
	return e.message
}

// Función que lanza un error personalizado manualmente

func lanzarErrorPersonalizado() error {
	// Creamos una instancia de nuestro error personalizado
	err := ErrorPersonalizado{"Esto es un error personalizado"}
	// Lo retornamos
	return err
}

func excepcionProc(a, b int) error {

	// Simulación de error personalizado

	if a < 0 {
		return errors.New("el primer parámetro no puede ser negativo")
	}

	// Simulación de error por división por cero

	if b == 0 {
		return errors.New("división por cero")
	}

	// Simulación de error genérico

	if b > 10 {
		return fmt.Errorf("el segundo parámetro es demasiado grande: %d", b)
	}

	// Si no hay error, imprimimos que la ejecución ha finalizado

	fmt.Println("Procesamiento completado sin errores")
	return nil
}

func practica10()/*main*/ {

	// Llamamos a la función que lanza el error personalizado
	
	err := lanzarErrorPersonalizado()
	// Capturamos el error
	if err != nil {
		// Imprimimos el error
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No se ha producido ningún error")
	}
	fmt.Println("La ejecución ha finalizado")

	// Llamamos a la función procesarParametros y capturamos los errores

	err = excepcionProc(5, 0)
	if err != nil {
		// Si hay un error, imprimimos el tipo de error
		fmt.Println("Error:", err)
	} else {
		// Si no hay error, imprimimos que la ejecución ha finalizado
		fmt.Println("No se ha producido ningún error")
	}
	fmt.Println("La ejecución ha finalizado")

}


