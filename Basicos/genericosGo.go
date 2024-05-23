package main

import "fmt"

func genericosGo() /*main*/ {
	
	fmt.Println(maxInt(1,3))     
	fmt.Println(maxString("a", "b"))
	fmt.Println(max("a","b"))

	fmt.Println(sum("key1", 1, 3))
	fmt.Println(sum("key2", 3, 3))
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

func max[T int | string](a, b T) T {
	if  a > b {
		return a
	}
	return b
}

// Se puede usar con types

type Numeric interface {
	int | float64 |float32 | int32 | int64
}

func sum[K string, T Numeric](key K, a,b T)T {
	fmt.Println(key)
	return a + b
}