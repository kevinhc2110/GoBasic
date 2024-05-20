package main

import "fmt"

// Se utilizan para manipular y gestionar la memoria de forma eficiente. Son útiles cuando necesitas pasar grandes estructuras de datos a funciones sin tener que copiarlas, cuando quieres modificar el valor de una variable en una función y que los cambios sean visibles fuera de esa función, o cuando necesitas trabajar con estructuras de datos enlazadas o implementar ciertas estructuras de datos complejas.
// Los punteros son variables que almacenan la dirección de memoria de otra variable. En lugar de contener un valor directamente, almacenan la ubicación en memoria de ese valor.

func punteros() /*main*/ {
    var entero int = 42
    var puntero *int = &entero
    fmt.Println(*puntero)
}

