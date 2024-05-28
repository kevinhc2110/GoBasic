package main

import "fmt"

func mario() {

	var altura int
	
	fmt.Println("Ingrese un tamaño del 1 al 8")
	fmt.Scanln(&altura)

	if altura >= 1 && altura <=8 {
		piramide(altura)
	}else {
		fmt.Println("Ingrese una altura validad")
	}
}

func piramide(altura int) {

	//  Ciclo principal el cual para al momento que i es 8

	for i := 1; i <= altura; i++ {

		// Ciclo para los espacios necesarios en el primer ciclo genera 7 espacios
		// En la primera interacción altura - i = 7
		
		for j := 0; j < altura-i; j++ {
			fmt.Print(" ")
		}

		// Ciclo para generar los asteriscos 
		// Solo se genera un # pues i en esta interacción el igual a 1

		for j := 0; j < i; j++ {
			fmt.Print("#")
		}

		// En este punto estamos general un espacio a la derecha durante cada interacción

		fmt.Print(" ")

		// Aquí generamos la otra pirámide después del espacio

		for j := 0; j < i; j++ {
			fmt.Print("#")
		}

		// Aquí hacemos salto de linea durante cada interacción

		println()
	}
}