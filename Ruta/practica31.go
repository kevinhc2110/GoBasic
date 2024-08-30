package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Estructura que representa un participante con su nombre y país
type Participante struct {
	Nombre string
	Pais   string
}

// Estructura que representa un evento con su nombre, lista de participantes y ganadores
type Evento struct {
	Nombre        string
	Participantes []Participante
	Ganadores     []Participante
}

// Estructura que mantiene el conteo de medallas (oro, plata y bronce) por país
type Medallero struct {
	Oro    int
	Plata  int
	Bronce int
}

// Lista global de eventos
var eventos []Evento

// Mapa para rastrear las medallas de cada país
var medallero = make(map[string]Medallero)

// Función para registrar un evento deportivo
func registrarEvento() {
	var nombreEvento string
	// Se solicita al usuario que ingrese el nombre del evento
	fmt.Print("Introduce el nombre del evento deportivo: ")
	fmt.Scanln(&nombreEvento)

	// Se crea un nuevo evento y se agrega a la lista de eventos
	evento := Evento{Nombre: nombreEvento}
	eventos = append(eventos, evento)
	fmt.Println("Evento registrado:", nombreEvento)
}

// Función para registrar un participante en un evento
func registrarParticipante() {
	// Si no hay eventos registrados, se informa al usuario
	if len(eventos) == 0 {
		fmt.Println("Primero debes registrar un evento.")
		return
	}

	var nombre, pais string
	// Se solicita al usuario que ingrese el nombre y país del participante
	fmt.Print("Introduce el nombre del participante: ")
	fmt.Scanln(&nombre)
	fmt.Print("Introduce el país del participante: ")
	fmt.Scanln(&pais)

	// Se crea un nuevo participante
	participante := Participante{Nombre: nombre, Pais: pais}

	// Se muestra la lista de eventos disponibles para seleccionar a cuál inscribir el participante
	for i, evento := range eventos {
		fmt.Printf("[%d] %s\n", i+1, evento.Nombre) // i+1 para mostrar los índices a partir de 1 (Por estética)
	}
	fmt.Print("Selecciona el evento al que quieres añadir el participante: ")
	var opcion int
	fmt.Scanln(&opcion)

	// Se añade el participante al evento seleccionado
	eventos[opcion-1].Participantes = append(eventos[opcion-1].Participantes, participante)
	fmt.Println("Participante registrado:", nombre, "de", pais)
}

// Función para simular los eventos y asignar medallas
func simularEventos() {
	// Si no hay eventos, se informa al usuario
	if len(eventos) == 0 {
		fmt.Println("No hay eventos registrados.")
		return
	}

	// Crear un generador de números aleatorios local
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Se recorre la lista de eventos para simularlos
	for i, evento := range eventos {
		// Si el evento tiene menos de 3 participantes, no se puede simular
		if len(evento.Participantes) < 3 {
			fmt.Println("El evento", evento.Nombre, "no tiene suficientes participantes.")
			continue
		}

		fmt.Println("Simulando evento:", evento.Nombre)

		// Se mezcla aleatoriamente la lista de participantes
		r.Shuffle(len(evento.Participantes), func(i, j int) {
			evento.Participantes[i], evento.Participantes[j] = evento.Participantes[j], evento.Participantes[i]
		})

		// Se asignan las medallas a los tres primeros participantes
		oro, plata, bronce := evento.Participantes[0], evento.Participantes[1], evento.Participantes[2]
		eventos[i].Ganadores = []Participante{oro, plata, bronce}

		// Se actualiza el medallero para cada país ganador
		medallero[oro.Pais] = sumarMedalla(medallero[oro.Pais], "oro")
		medallero[plata.Pais] = sumarMedalla(medallero[plata.Pais], "plata")
		medallero[bronce.Pais] = sumarMedalla(medallero[bronce.Pais], "bronce")

		// Se muestran los ganadores del evento
		fmt.Printf("Ganadores de %s:\nOro: %s (%s)\nPlata: %s (%s)\nBronce: %s (%s)\n\n",
			evento.Nombre, oro.Nombre, oro.Pais, plata.Nombre, plata.Pais, bronce.Nombre, bronce.Pais)
	}
}

// Función que actualiza el conteo de medallas para un país
func sumarMedalla(m Medallero, tipo string) Medallero {
	// Dependiendo del tipo de medalla, se incrementa el respectivo contador
	switch tipo {
	case "oro":
		m.Oro++
	case "plata":
		m.Plata++
	case "bronce":
		m.Bronce++
	}
	return m
}

// Función para generar un informe final de los eventos y el ranking de países
func generarInforme() {
	// Si no se han simulado eventos, se informa al usuario
	if len(eventos) == 0 {
		fmt.Println("No hay eventos simulados.")
		return
	}

	// Se muestran los ganadores de cada evento
	fmt.Println("\n--- Informe Final ---")
	for _, evento := range eventos {
		if len(evento.Ganadores) > 0 {
			fmt.Printf("Evento: %s\n", evento.Nombre)
			fmt.Printf("Oro: %s (%s)\n", evento.Ganadores[0].Nombre, evento.Ganadores[0].Pais)
			fmt.Printf("Plata: %s (%s)\n", evento.Ganadores[1].Nombre, evento.Ganadores[1].Pais)
			fmt.Printf("Bronce: %s (%s)\n", evento.Ganadores[2].Nombre, evento.Ganadores[2].Pais)
			fmt.Println()
		}
	}

	// Se muestra el ranking de países según el número de medallas
	fmt.Println("\n--- Ranking de países por medallas ---")
	var paises []string
	for pais := range medallero {
		paises = append(paises, pais)
	}

	// Se ordenan los países en función del número de medallas de oro
	sort.Slice(paises, func(i, j int) bool {
		return medallero[paises[i]].Oro > medallero[paises[j]].Oro
	})

	// Se muestran los países con su respectivo conteo de medallas
	for _, pais := range paises {
		m := medallero[pais]
		fmt.Printf("%s -> Oro: %d, Plata: %d, Bronce: %d\n", pais, m.Oro, m.Plata, m.Bronce)
	}
}

// Función principal que muestra el menú de opciones y ejecuta las acciones
func practica31() {
	for {
		fmt.Println("\n--- Menú ---")
		fmt.Println("1. Registrar evento deportivo")
		fmt.Println("2. Registrar participante")
		fmt.Println("3. Simular eventos")
		fmt.Println("4. Generar informe")
		fmt.Println("5. Salir")
		fmt.Print("Selecciona una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		// Se ejecuta la acción correspondiente según la opción seleccionada por el usuario
		switch opcion {
		case 1:
			registrarEvento() // Registrar un nuevo evento
		case 2:
			registrarParticipante() // Registrar un participante en un evento
		case 3:
			simularEventos() // Simular los eventos registrados
		case 4:
			generarInforme() // Generar el informe final de eventos y medallas
		case 5:
			fmt.Println("¡Hasta luego!") // Salir del programa
			return
		default:
			fmt.Println("Opción no válida. Inténtalo de nuevo.")
		}
	}
}
