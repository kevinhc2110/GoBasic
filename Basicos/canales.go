package main

import "fmt"

// Se utilizan para la comunicación entre goroutines, especialmente en el contexto de la concurrencia y la programación concurrente. Son útiles cuando necesitas coordinar la ejecución de diferentes partes de tu programa y compartir datos de forma segura entre ellas.
// Los canales son una característica clave en Go para la comunicación y la sincronización entre goroutines (hilos de ejecución concurrentes). Permiten que las goroutines se comuniquen entre sí enviando y recibiendo valores de manera segura.

func canales() /*main*/ {
    canal := make(chan int)
    go func() {
        canal <- 42
    }()
    resultado := <-canal
    fmt.Println(resultado)
}