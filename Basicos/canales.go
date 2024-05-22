package main

import (
	"errors"
	"fmt"
	"time"
)

// Se utilizan para la comunicación entre goroutines, especialmente en el contexto de la concurrencia y la programación concurrente. Son útiles cuando necesitas coordinar la ejecución de diferentes partes de tu programa y compartir datos de forma segura entre ellas.
// Los canales son una característica clave en Go para la comunicación y la sincronización entre goroutines (hilos de ejecución concurrentes). Permiten que las goroutines se comuniquen entre sí enviando y recibiendo valores de manera segura.

func canales() /*main*/ {
    canal := make(chan int)
    go func() {
        canal <- 42
    }()
    resultado := <-canal
    fmt.Println(resultado)

    // Mas tipos de canales

    // Canales Sin Buffer (Unbuffered Channels):
    // La comunicación en un canal sin buffer se realiza de manera sincronizada.

    ch := make(chan int)
    go func() {
        
    }()
    val := <-ch
    fmt.Println(val)

    //Canales con Buffer (Buffered Channels):
    // La comunicación se realiza de manera asíncrona hasta que el buffer se llena.

    ch1 := make(chan int, 2)
    ch1 <- 1
    ch1 <- 2
    val1 := <-ch
    val2 := <-ch
    fmt.Println(val1, val2)

    // Canales Unidireccionales (One-Way Channels):
    // Los canales unidireccionales permiten restringir el uso de un canal solo para enviar o solo para recibir.

    ch2 := make(chan int)

    var sendCh chan<- int = ch2 // Solo para enviar
    var recvCh <-chan int = ch2 // Solo para recibir

    go func() {
        sendCh <- 1
    }()
    val3 := <-recvCh
    fmt.Println(val3)

    // Select Statements
    // La declaración select permite esperar en múltiples operaciones de canal. Es útil para manejar varios canales al mismo tiempo y realizar la operación que esté lista primero.

    ch3 := make(chan int)
    ch4 := make(chan int)

    go func() {
        ch3 <- 1
    }()
    go func() {
        ch4 <- 2
    }()

    select {
    case val1 := <-ch3:
        fmt.Println("Received from ch1:", val1)
    case val2 := <-ch4:
        fmt.Println("Received from ch2:", val2)
    }

    // Cerrar Canales (Closing Channels):
    // Los canales pueden cerrarse para indicar que no se enviarán más datos. Los receptores pueden comprobar si un canal está cerrado.

    ch5 := make(chan int)
    go func() {
        for i := 0; i < 3; i++ {
        ch5 <- i
        }
        close(ch5)
    }()

    for val := range ch5 {
        fmt.Println(val)
    }

    // Rango en Canales (Range Over Channels):
    // Se puede usar un bucle range para recibir valores de un canal hasta que se cierre.

    ch6 := make(chan int)
    go func() {
        for i := 0; i < 3; i++ {
            ch6 <- i
        }
        close(ch6)
    }()

    for val := range ch6 {
        fmt.Println(val)
    }

    // Canal de Señalización (done channel):
    // Utilizado para notificar que una Goroutine ha terminado su trabajo.

    done := make(chan bool)
    go func() {
        // Hacer algo
        done <- true
    }()
    <-done

    // Canal de Errores:
    // Utilizado para manejar errores de forma concurrente.

    errCh := make(chan error)
    go func() {
        err := someFunction()
        errCh <- err
    }()
    err := <-errCh
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Canales de Datos y Control:
    // A veces se utilizan dos canales: uno para los datos y otro para el control.

    dataCh := make(chan int)
    controlCh := make(chan struct{})
    go func() {
        for i := 0; i < 5; i++ {
            dataCh <- i
        }
        controlCh <- struct{}{}
    }()
    go func() {
        for {
            select {
            case data := <-dataCh:
                fmt.Println("Received data:", data)
            case <-controlCh:
                fmt.Println("Control signal received, exiting.")
                return
            }
        }
    }()

    // Canal de Resultado:
    // Utilizado para devolver resultados de funciones concurrentes.

    resultCh := make(chan int)
    go func() {
        result := computeSomething()
        resultCh <- result
    }()
    result := <-resultCh
    fmt.Println("Result:", result)
}

func someFunction() error {
	time.Sleep(2 * time.Second) // Simulando un trabajo que toma tiempo
	return errors.New("simulated error") // Devolver un error simulado
}

func computeSomething() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		time.Sleep(100 * time.Millisecond) // Simulando un trabajo que toma tiempo
	}
	return sum
}




