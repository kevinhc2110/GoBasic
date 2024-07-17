//* Funciones y alcance

package main

import "fmt"

/*
DIFICULTAD EXTRA:
Crea una función que reciba dos parámetros de tipo cadena de texto y retorne un número.
- La función imprime todos los números del 1 al 100. Teniendo en cuenta que:
- Si el número es múltiplo de 3, muestra la cadena de texto del primer parámetro.
- Si el número es múltiplo de 5, muestra la cadena de texto del segundo parámetro.
- Si el número es múltiplo de 3 y de 5, muestra las dos cadenas de texto concatenadas.
- La función retorna el número de veces que se ha impreso el número en lugar de los textos.
Presta especial atención a la sintaxis que debes utilizar en cada uno de los casos.
Cada lenguaje sigue una convenciones que debes de respetar para que el código se entienda.
*/

func practica02() /*main*/ {
    fun02("Pedrito","Flores")
}

func fun02(cadena1 string, cadena2 string) int {
    var numero int

    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println(cadena1+cadena2)
        } else if i%3 == 0 {
            fmt.Println(cadena1)
        } else if i%5 == 0 {
            fmt.Println(cadena2)
        } else {
            fmt.Println(i)
            numero++
        }
    }

    fmt.Println("Cantidad de veces que salió un número:", numero)
    return numero
}