package main

import (
	"fmt"
	"sort"
	"time"
)

/* DIFICULTAD EXTRA:
Dada una lista de estudiantes (con sus nombres, fecha de nacimiento y
lista de calificaciones), utiliza funciones de orden superior para
realizar las siguientes operaciones de procesamiento y análisis:
- Promedio calificaciones: Obtiene una lista de estudiantes por nombre
  y promedio de sus calificaciones.
- Mejores estudiantes: Obtiene una lista con el nombre de los estudiantes
  que tienen calificaciones con un 9 o más de promedio.
- Nacimiento: Obtiene una lista de estudiantes ordenada desde el más joven.
- Mayor calificación: Obtiene la calificación más alta de entre todas las
  de los alumnos.
- Una calificación debe estar comprendida entre 0 y 10 (admite decimales).
*/

type Estudiante struct{

	Nombres string
	FechaNacimiento time.Time
	Calificaciones []float64

}

func main() {

	estudiantes := []Estudiante{
		{	
		Nombres:          "Juan Pérez",
		FechaNacimiento: time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC),
		Calificaciones:  []float64{6.5, 7.0, 8.0},
		},
		{
			Nombres:          "María López",
			FechaNacimiento: time.Date(1999, time.March, 22, 0, 0, 0, 0, time.UTC),
			Calificaciones:  []float64{9.2, 9.5, 9.7},
		},
		{
			Nombres:          "Roberto Zuluaga",
			FechaNacimiento: time.Date(1997, time.December, 29, 0, 0, 0, 0, time.UTC),
			Calificaciones:  []float64{8.5, 9.0, 7.0},
		},
	}

	// Obtener promedio de calificaciones por estudiante
	promedios := PromedioCalificaciones(estudiantes)
	fmt.Println("Promedio de calificaciones por estudiante:")
	for nombre, promedio := range promedios {
		fmt.Printf("%s: %.2f\n", nombre, promedio)
	}

	// Obtener la lista de mejores estudiantes
	mejores := MejoresEstudiantes(estudiantes)
	fmt.Println("\nMejores estudiantes (promedio >= 9):")
	for _, nombre := range mejores {
		fmt.Println(nombre)
	}

	// Obtener la lista de estudiantes ordenada por fecha de nacimiento
	ordenados := OrdenarPorNacimiento(estudiantes)
	fmt.Println("\nEstudiantes ordenados por fecha de nacimiento (más joven primero):")
	for _, estudiante := range ordenados {
		fmt.Printf("%s: %s\n", estudiante.Nombres, estudiante.FechaNacimiento.Format("02/01/2006"))
	}

	// Obtener la mayor calificación entre todos los estudiantes
	maxCalificacion := MayorCalificacion(estudiantes)
	fmt.Printf("\nMayor calificación entre todos los estudiantes: %.2f\n", maxCalificacion)
}

func (e* Estudiante) PromedioCalificaciones() float64  {

	if len(e.Calificaciones) == 0 {
		return 0.0
	}
	suma := 0.0

	for _, calificacion := range e.Calificaciones {
		suma += calificacion
	}
	return suma/ float64(len(e.Calificaciones))
	
}

// Lista promedio de califaciones

func PromedioCalificaciones(estudiantes []Estudiante) map[string]float64 {
	promedios := make(map[string]float64)
	for _, estudiante := range estudiantes {
		promedios[estudiante.Nombres] = estudiante.PromedioCalificaciones()
	}
	return promedios
}

// Mejores estudiantes 

func MejoresEstudiantes(estudiantes []Estudiante) []string {
	mejores := []string{}
	for _, estudiante := range estudiantes {
		if estudiante.PromedioCalificaciones() >= 9.0 {
			mejores = append(mejores, estudiante.Nombres)
		}
	}
	return mejores
}

// Nacimiento (ordenado desde el más joven)

func OrdenarPorNacimiento(estudiantes []Estudiante) []Estudiante {
	ordenados := make([]Estudiante, len(estudiantes))
	copy(ordenados, estudiantes)
	sort.Slice(ordenados, func(i, j int) bool {
		return ordenados[i].FechaNacimiento.After(ordenados[j].FechaNacimiento)
	})
	return ordenados
}

// Mayor calificación:

func MayorCalificacion(estudiantes []Estudiante) float64 {
	maxCalificacion := 0.0
	for _, estudiante := range estudiantes {
		for _, calificacion := range estudiante.Calificaciones {
			if calificacion > maxCalificacion {
				maxCalificacion = calificacion
			}
		}
	}
	return maxCalificacion
}
