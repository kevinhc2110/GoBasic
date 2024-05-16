package main

import "fmt"

// Similar a los arrays pero con tamaño dinámico

func slices() {

	// Creación de un slice vacío
	
	var s []int

	// Agregar elementos al slice usando la función append

	s = append(s, 2, 3, 4)

	fmt.Println(s)

	 // Acceder a elementos individuales del slice

	 fmt.Println("Primer elemento:", s[0])
	 fmt.Println("Segundo elemento:", s[1])

	 // Modificar elementos del slice

	 s[0] = 5
	 fmt.Println("Slice modificado:", s)

	 // Obtener longitud y capacidad del slice

	 fmt.Println("Longitud del slice:", len(s))
	 fmt.Println("Capacidad del slice:", cap(s))

	 // Crear un slice con make especificando longitud y capacidad

	 s2 := make([]int, 3, 5) // longitud: 3, capacidad: 5
	 fmt.Println("Slice s2:", s2)
	 fmt.Println("Longitud de s2:", len(s2))
	 fmt.Println("Capacidad de s2:", cap(s2))

	 // Iterar sobre los elementos del slice

	 fmt.Println("Iterando sobre los elementos de s:")
	 for i, v := range s {
			 fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	 }
}