package main

import "fmt"

func interaciones() {

    // Utilizando un bucle for

    fmt.Println("Imprimiendo números del 1 al 10 utilizando un bucle for:")
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Utilizando un bucle for con la función range

    fmt.Println("\nImprimiendo números del 1 al 10 utilizando un bucle for con la función range:")
    for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
        fmt.Println(num)
    }

    // Utilizando un bucle for con una variable de control y condicional if
		
    fmt.Println("\nImprimiendo números del 1 al 10 utilizando un bucle for con una variable de control y condicional if:")
    i := 1
    for {
        fmt.Println(i)
        if i == 10 {
            break
        }
        i++
    }
}
