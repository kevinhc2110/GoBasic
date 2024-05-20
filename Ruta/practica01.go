package main

import "fmt"

func practica01() /*main*/ {

	// https://go.dev/doc/

	for i := 10; i <= 55; i++ {

		if (i%2 == 0 && i != 16 && i%3 != 0) {
			fmt.Println(i)
		}
	}
	
}
