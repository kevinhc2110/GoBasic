package main

import "fmt"

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