package main

import (
	"container/list"
	"fmt"
)

func /*estruControl*/ main() {

	num1 := 1
	name := "Hola"

	if num1 == 5 {
		fmt.Println("El valor es 5")
	} else if num1 == 10 {
		fmt.Println("El valor es 10")
	} else {
		fmt.Println("El valor no es 5")
	}

	// And
	if (num1 == 7 && name == "Coco"){
		fmt.Println("El valor es 5")
	} else if (num1 == 7 || name == "Pepe"){ // Or
		fmt.Println("El valor no es 5")
	}

	// Array

	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 1
	myArray[2] = 1

	fmt.Println(myArray)
	fmt.Println(myArray[2])

	// Map

	myMap := make(map[string]int)
	myMap["Brais"] = 36
	myMap["Kevin"] = 26

	fmt.Println(myMap)
	fmt.Println(myMap["Brais"])

	myMap2 := map[string]int{"Brais":36, "Kevin":26}
	fmt.Println(myMap2)

	// Lista

	myList:= list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	fmt.Println(myList) //Representación de punteros
	fmt.Println(myList.Back().Value) // Representa el ultimo valor de la lista

	// Bucles

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for key, value := range myMap {
		fmt.Println(key,value)
	}

	for index, value := range myArray {
		fmt.Println(index,value)
	}


	// Switch

	option:= 1

	switch option {
		case 1:
			fmt.Println("Caso 1")
	 	case 2:
      fmt.Println("Caso 2")
    default:
      fmt.Println("Ningún caso")
	}

	// Defer: Pospone la ejecución de una función hasta que la función que la contiene haya terminado

	defer fmt.Println("Esto se ejecuta al final")
    fmt.Println("Esto se ejecuta primero")

	// Panic: Detiene la ejecución normal de una función. Puede utilizarse para lanzar errores.

	panic("Se produjo un error")


}

// Recover: Se utiliza para manejar una situación de pánico y continuar la ejecución.	 

func safeDivide(a, b int) (result int) {
	defer func() {
			if r := recover(); r != nil {
					fmt.Println("Recuperado del pánico:", r)
					result = 0
			}
	}()
	if b == 0 {
			panic("división por cero")
	}
	return a / b
}

// No hay necesidad de colocar el panic porque implícitamente esta en el error

func safeDivide1(a, b int) (result int) {
	defer func() {
			if r := recover(); r != nil {
					fmt.Println("Recuperado del pánico:", r)
					result = 0
			}
	}()
	return a / b
}





	

