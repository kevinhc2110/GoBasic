package main

import "fmt"

/*
EJERCICIO:
Explora el "Principio SOLID de Sustitución de Liskov (Liskov Substitution Principle, LSP)"
y crea un ejemplo simple donde se muestre su funcionamiento
de forma correcta e incorrecta.
DIFICULTAD EXTRA (opcional):
Crea una jerarquía de vehículos. Todos ellos deben poder acelerar y frenar, así como
cumplir el LSP.
Instrucciones:
1. Crea la clase Vehículo.
2. Añade tres subclases de Vehículo.
3. Implementa las operaciones "acelerar" y "frenar" como corresponda.
4. Desarrolla un código que compruebe que se cumple el LSP.
*/

// Vehiculo es una interfaz que representa un vehículo genérico
type Vehiculo interface {
	Acelerar() string
	Frenar() string
}

// Coche es una estructura que representa un coche
type Coche struct{}

func (c Coche) Acelerar() string {
	return "El coche está acelerando"
}

func (c Coche) Frenar() string {
	return "El coche está frenando"
}

// Motocicleta es una estructura que representa una motocicleta
type Motocicleta struct{}

func (m Motocicleta) Acelerar() string {
	return "La motocicleta está acelerando"
}

func (m Motocicleta) Frenar() string {
	return "La motocicleta está frenando"
}

// Bicicleta es una estructura que representa una bicicleta
type Bicicleta struct{}

func (b Bicicleta) Acelerar() string {
	return "La bicicleta está acelerando"
}

func (b Bicicleta) Frenar() string {
	return "La bicicleta está frenando"
}

func practica28() {
	// Creamos una lista de vehículos
	vehiculos := []Vehiculo{
		Coche{},
		Motocicleta{},
		Bicicleta{},
	}

	// Probamos acelerar y frenar con cada vehículo
	for _, vehiculo := range vehiculos {
		fmt.Println(vehiculo.Acelerar())
		fmt.Println(vehiculo.Frenar())
	}
}
