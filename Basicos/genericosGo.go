package main

import "fmt"

func genericosGo() /*main*/ {
	
	fmt.Println(maxInt(2,3))     
	fmt.Println(maxString("a", "b"))
	
	
}

// Considere una función que compara dos valores y devuelve el valor más grande. En Go tradicional, se necesitarían dos funciones separadas para comparar números enteros y cadenas de texto.

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxString(a, b string) string {
	if a > b {
		return a
	}
	return b
}

// Con datos genéricos, se puede escribir una única función genérica que compara dos valores de cualquier tipo comparable.

// func max[T comparable](a, b T) T {
// 	if  a > b {
// 		return a
// 	}
// 	return b
// }