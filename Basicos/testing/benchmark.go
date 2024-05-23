package main

import "testing"

func BenchmarkSumar(b *testing.B) {
	
   var result int
    for i := 0; i < b.N; i++ {
        result = Sumar(3, 7)
    }
    // Usar el resultado para evitar la advertencia SA4017
    _ = result
}