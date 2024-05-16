package main

import (
	"fmt"
	"reflect"
)

//Declaración de variables
func variables /*main*/()  {
	var num int = 3
	var name string = "KevinDev"

	fmt.Println(num, name)

	// Declara e inicializa múltiples variables
	
	var a, b, c int = 1, 2, 3  
  
	fmt.Println(a, b, c)

	num = 0
	name = "Sofia"

	var num2 = 0;
	var name2 string = "Dev"

	fmt.Println(num2, name2)

	//num2 = "" Seria un error porque el tipo que tomo anteriormente fue de int

	num2 = num2 + 8
	fmt.Println(num2)

	fmt.Println(num2 - 1)
	
	//Concatenar tipos diferentes

	fmt.Println(name2, num2)

	//Mostrar el tipo por consola

	fmt.Println(reflect.TypeOf(num))

	//Sumar entero y float

	var float1 float64 = 6.5
	fmt.Println(float1 + float64(num2))

	//Mas tipos de datos

	var bool1 bool = false
	fmt.Println(bool1)
 
	//Evitar usar el var para declarar e inicializar

	name3 := "Papi";
	fmt.Println(name3)

	//Constantes
	const myConst = "Kevin Llora"
}


