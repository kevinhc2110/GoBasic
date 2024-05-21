package main

import "testing"

// TestSumar verifica que la función Sumar retorne el resultado correcto.
// Define un test TestSumar que utiliza el paquete testing.
// Crea una lista de casos de prueba con entradas a, b y el resultado esperado expected.
// Itera sobre los casos de prueba y llama a la función Sumar con las entradas.
// Verifica si el resultado coincide con el esperado. Si no es así, llama a t.Errorf para reportar un error con detalles del caso fallido.

func TestSumar(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{0, 0, 0},
		{-1, -1, -2},
		{-1, 1, 0},
	}

	for _, tt := range tests {
		result := Sumar(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Sumar(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}