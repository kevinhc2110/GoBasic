package main

import "fmt"

func estruClase() /*main*/ {
	// No hay clases si no estructuras
	// Estructuras

	type MyStruct struct {
		name string
		age int
	}

	myStruct := MyStruct {"Brais", 36}
	fmt.Println(myStruct)

}


