package main

import (
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}

func excepciones() /*main*/ {

	// Forzando un error al dividir por cero
	
	resultado, err := division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Resultado de la división:", resultado)

}