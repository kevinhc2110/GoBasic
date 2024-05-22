package main

/*
DIFICULTAD EXTRA:
Escribe el mayor número de mecanismos que posea tu lenguaje
para iterar valores. ¿Eres capaz de utilizar 5? ¿Y 10?
*/

import "fmt"

func practica17() /*main*/ {

    // Utilizando un bucle for clásico

    fmt.Println("Bucle for clásico:")
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Utilizando un bucle for con la función range para un slice

    fmt.Println("\nBucle for con la función range para un slice:")
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for _, num := range slice {
        fmt.Println(num)
    }

    // Utilizando un bucle for con la función range para un map

    fmt.Println("\nBucle for con la función range para un map:")
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    for key, value := range m {
        fmt.Printf("%s: %d\n", key, value)
    }

    // Utilizando un bucle for con una condición al principio

    fmt.Println("\nBucle for con una condición al principio:")
    j := 1
    for j <= 10 {
        fmt.Println(j)
        j++
    }

    // Utilizando un bucle for sin condición (bucle infinito con break)

    fmt.Println("\nBucle for sin condición (bucle infinito con break):")
    k := 1
    for {
        fmt.Println(k)
        if k == 10 {
            break
        }
        k++
    }

    // Utilizando un bucle for con una condición al final (do-while)

    fmt.Println("\nBucle for con una condición al final (do-while):")
    l := 1
    for {
        fmt.Println(l)
        l++
        if l > 10 {
            break
        }
    }

    // Utilizando un bucle for con un label (goto)

    fmt.Println("\nBucle for con un label (goto):")
    s := 1
		OuterLoop:
    for s <= 10 {
        fmt.Println(m)
        s++
        goto OuterLoop
    }

    // Utilizando un bucle for con un slice y puntero al slice

    fmt.Println("\nBucle for con un slice y puntero al slice:")
    slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    ptr := &slice2
    for i := range *ptr {
        fmt.Println((*ptr)[i])
    }

    // Utilizando un bucle for con un string

    fmt.Println("\nBucle for con un string:")
    str := "Hello, world!"
    for i := 0; i < len(str); i++ {
        fmt.Println(string(str[i]))
    }

    // Utilizando un bucle for con un channel (select)
		// Se crea un canal de enteros ch y se inicia una goroutine que envía los números del 1 al 10 al canal. Luego, en el bucle for, se utiliza la sintaxis range para iterar sobre los valores enviados al canal y se imprime cada número. 
		
    fmt.Println("\nBucle for con un channel (select):")
    ch := make(chan int)
    go func() {
        for i := 1; i <= 10; i++ {
            ch <- i
        }
        close(ch)
    }()
    for num := range ch {
        fmt.Println(num)
    }
}
