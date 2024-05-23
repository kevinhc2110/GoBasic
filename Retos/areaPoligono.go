package main

import "fmt"

/*
Crea una única función (importante que sólo sea una) que sea capaz
de calcular y retornar el área de un polígono.
- La función recibirá por parámetro sólo UN polígono a la vez.
- Los polígonos soportados serán Triángulo, Cuadrado y Rectángulo.
- Imprime el cálculo del área de un polígono de cada tipo.
*/

func areaPoligono() {

	// Opción 1

	fmt.Printf("Área del triángulo: %.2f\n", CalcularArea("triangulo", 5, 3))
	fmt.Printf("Área del cuadrado: %.2f\n", CalcularArea("cuadrado", 4))
	fmt.Printf("Área del rectángulo: %.2f\n", CalcularArea("rectangulo", 6, 2))

	// Opción 2

	triangulo := Triangulo{Base: 5, Altura: 3}
	cuadrado := Cuadrado{Lado: 4}
	rectangulo := Rectangulo{Base: 6, Altura: 2}
	
	fmt.Printf("Área del triángulo: %.2f\n", CalcularArea1(triangulo))
	fmt.Printf("Área del cuadrado: %.2f\n", CalcularArea1(cuadrado))
	fmt.Printf("Área del rectángulo: %.2f\n", CalcularArea1(rectangulo))

	// Opción 3

	baseTriangulo, alturaTriangulo := 5.0, 3.0
	ladoCuadrado := 4.0
	baseRectangulo, alturaRectangulo := 6.0, 2.0

	fmt.Printf("Área del triángulo: %.2f\n", AreaTriangulo(baseTriangulo, alturaTriangulo))
	fmt.Printf("Área del cuadrado: %.2f\n", AreaCuadrado(ladoCuadrado))
	fmt.Printf("Área del rectángulo: %.2f\n", AreaRectangulo(baseRectangulo, alturaRectangulo))
}

// Función que calcula y retorna el área de un polígono
// La sintaxis params ...float64 en la firma de la función CalcularArea indica que la función puede aceptar un número variable de parámetros de tipo float64.

func CalcularArea(tipo string, params ...float64) float64 {
	
	switch tipo {
	case "triangulo":
		base, altura := params[0], params[1]
		return (base * altura) / 2
	case "cuadrado":
		lado := params[0]
		return lado * lado
	case "rectangulo":
		base, altura := params[0], params[1]
		return base * altura
	default:
		fmt.Println("Tipo de polígono no soportado")
		return 0
	}
}

// Opción 2

// Definición de la interfaz Polígono

type Poligono interface {
	Area() float64
}

// Definición de la estructura Triangulo

type Triangulo struct {
	Base   float64
	Altura float64
}

// Implementación del método Area() para Triangulo

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

// Definición de la estructura Cuadrado

type Cuadrado struct {
	Lado float64
}

// Implementación del método Area() para Cuadrado

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

// Definición de la estructura Rectángulo

type Rectangulo struct {
	Base   float64
	Altura float64
}

// Implementación del método Area() para Rectángulo

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

// Función que calcula y retorna el área de un polígono

func CalcularArea1(poligono Poligono) float64 {
	return poligono.Area()
}

// Opción 3

// Función que calcula y retorna el área de un triángulo

func AreaTriangulo(base, altura float64) float64 {
	return (base * altura) / 2
}

// Función que calcula y retorna el área de un cuadrado

func AreaCuadrado(lado float64) float64 {
	return lado * lado
}

// Función que calcula y retorna el área de un rectángulo

func AreaRectangulo(base, altura float64) float64 {
	return base * altura
}

