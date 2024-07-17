//* Conjuntos

package main

import (
	"fmt"
)

func practica18() /*main*/ {

	// Definir conjuntos

	conjunto1 := []int{1, 2, 3, 4, 5}
	conjunto2 := []int{4, 5, 6, 7, 8}

	// Unión

	union := union(conjunto1, conjunto2)
	fmt.Println("Unión:", union)

	// Intersección

	interseccion := interseccion(conjunto1, conjunto2)
	fmt.Println("Intersección:", interseccion)

	// Diferencia

	diferencia := diferencia(conjunto1, conjunto2)
	fmt.Println("Diferencia (conjunto1 - conjunto2):", diferencia)

	// Diferencia simétrica

	diferenciaSimetrica := diferenciaSimetrica(conjunto1, conjunto2)
	fmt.Println("Diferencia simétrica:", diferenciaSimetrica)
	
}

// Función para calcular la unión de dos conjuntos

func union(conjunto1, conjunto2 []int) []int {

	// Crear un mapa para almacenar los elementos únicos

	unique := make(map[int]bool)
	union := []int{}

	// Agregar elementos del primer conjunto al mapa

	for _, elemento := range conjunto1 {
		unique[elemento] = true
	}

	// Agregar elementos del segundo conjunto al mapa

	for _, elemento := range conjunto2 {
		unique[elemento] = true
	}

	// Extraer elementos únicos del mapa y agregarlos a la unión

	for elemento := range unique {
		union = append(union, elemento)
	}

	return union
}

// Función para calcular la intersección de dos conjuntos

func interseccion(conjunto1, conjunto2 []int) []int {

	// Crear un mapa para almacenar los elementos del primer conjunto

	elements := make(map[int]bool)
	interseccion := []int{}

	// Agregar elementos del primer conjunto al mapa

	for _, elemento := range conjunto1 {
		elements[elemento] = true
	}

	// Verificar si los elementos del segundo conjunto están en el mapa

	for _, elemento := range conjunto2 {
		if elements[elemento] {
			interseccion = append(interseccion, elemento)
		}
	}

	return interseccion
}

// Función para calcular la diferencia de dos conjuntos

func diferencia(conjunto1, conjunto2 []int) []int {
	diferencia := []int{}

	// Crear un mapa para almacenar los elementos del segundo conjunto

	elements := make(map[int]bool)
	for _, elemento := range conjunto2 {
		elements[elemento] = true
	}

	// Verificar si los elementos del primer conjunto están en el mapa

	for _, elemento := range conjunto1 {
		if !elements[elemento] {
			diferencia = append(diferencia, elemento)
		}
	}

	return diferencia
}

// Función para calcular la diferencia simétrica de dos conjuntos

func diferenciaSimetrica(conjunto1, conjunto2 []int) []int {

	diferenciaSimetrica := []int{}

	// Al asignar _ a diferenciaSimetrica estás indicando explícitamente que no tienes intención de usar el valor de diferenciaSimetrica en el resto del código
	_ = diferenciaSimetrica

	// Obtener la diferencia entre el primer y segundo conjunto

	dif1 := diferencia(conjunto1, conjunto2)

	// Obtener la diferencia entre el segundo y primer conjunto

	dif2 := diferencia(conjunto2, conjunto1)

	// Unir ambas diferencias

	diferenciaSimetrica = append(dif1, dif2...)

	return diferenciaSimetrica
}




