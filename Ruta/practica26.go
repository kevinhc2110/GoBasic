//* Solid: Principio de responsabilidad única (SRP)

package main

import (
	"fmt"
	"time"
)

/*
DIFICULTAD EXTRA:
Desarrolla un sistema de gestión para una biblioteca. El sistema necesita
manejar diferentes aspectos como el registro de libros, la gestión de usuarios
y el procesamiento de préstamos de libros.
Requisitos:
1. Registrar libros: El sistema debe permitir agregar nuevos libros con
información básica como título, autor y número de copias disponibles.
2. Registrar usuarios: El sistema debe permitir agregar nuevos usuarios con
información básica como nombre, número de identificación y correo electrónico.
3. Procesar préstamos de libros: El sistema debe permitir a los usuarios
tomar prestados y devolver libros.
Instrucciones:
1. Diseña una clase que no cumple el SRP: Crea una clase Library que maneje
los tres aspectos mencionados anteriormente (registro de libros, registro de
usuarios y procesamiento de préstamos).
2. Refactoriza el código: Separa las responsabilidades en diferentes clases
siguiendo el Principio de Responsabilidad Única.
*/

//*Clase que No Cumple el SRP*

type Libro struct {
	Titulo            string
	Autor             string
	CopiasDisponibles int
}

type Usuario struct {
	Nombre               string
	NumeroIdentificacion string
	Correo               string
}

type Prestamo struct {
	Usuario         Usuario
	Libro           Libro
	FechaPrestamo   time.Time
	FechaDevolucion time.Time
}

type Library struct {
	Libros    []Libro
	Usuarios  []Usuario
	Prestamos []Prestamo
}

func (l *Library) RegistrarLibro(titulo, autor string, copias int) {
	libro := Libro{Titulo: titulo, Autor: autor, CopiasDisponibles: copias}
	l.Libros = append(l.Libros, libro)
	fmt.Printf("Libro registrado: %s\n", titulo)
}

func (l *Library) RegistrarUsuario(nombre, id, correo string) {
	usuario := Usuario{Nombre: nombre, NumeroIdentificacion: id, Correo: correo}
	l.Usuarios = append(l.Usuarios, usuario)
	fmt.Printf("Usuario registrado: %s\n", nombre)
}

func (l *Library) ProcesarPrestamo(usuarioID, libroTitulo string) {
	for i, libro := range l.Libros {
		if libro.Titulo == libroTitulo && libro.CopiasDisponibles > 0 {
			for _, usuario := range l.Usuarios {
				if usuario.NumeroIdentificacion == usuarioID {
					prestamo := Prestamo{
						Usuario:       usuario,
						Libro:         libro,
						FechaPrestamo: time.Now(),
					}
					l.Prestamos = append(l.Prestamos, prestamo)
					l.Libros[i].CopiasDisponibles--
					fmt.Printf("Libro prestado: %s a %s\n", libroTitulo, usuario.Nombre)
					return
				}
			}
		}
	}
	fmt.Println("No se pudo procesar el préstamo")
}

// func nosrp() {
// 	biblioteca := Library{}
// 	biblioteca.RegistrarLibro("El Quijote", "Miguel de Cervantes", 3)
// 	biblioteca.RegistrarUsuario("Juan Perez", "1234", "juan@example.com")
// 	biblioteca.ProcesarPrestamo("1234", "El Quijote")
// }

// La clase Library maneja múltiples responsabilidades: registrar libros, registrar usuarios y procesar préstamos. Esto va en contra del SRP.

//*Refactorización del Código para Cumplir con el SRP*

// Libro representa un libro en la biblioteca
type Libro1 struct {
	Titulo            string
	Autor             string
	CopiasDisponibles int
}

// Usuario representa un usuario de la biblioteca
type Usuario1 struct {
	Nombre               string
	NumeroIdentificacion string
	Correo               string
}

// Prestamo representa un préstamo de libro
type Prestamo1 struct {
	Usuario         Usuario
	Libro           Libro
	FechaPrestamo   time.Time
	FechaDevolucion time.Time
}

// ManejadorLibros maneja la lógica relacionada con los libros
type ManejadorLibros struct {
	Libros []Libro
}

func (m *ManejadorLibros) RegistrarLibro(titulo, autor string, copias int) {
	libro := Libro{Titulo: titulo, Autor: autor, CopiasDisponibles: copias}
	m.Libros = append(m.Libros, libro)
	fmt.Printf("Libro registrado: %s\n", titulo)
}

func (m *ManejadorLibros) DisminuirCopia(titulo string) bool {
	for i, libro := range m.Libros {
		if libro.Titulo == titulo && libro.CopiasDisponibles > 0 {
			m.Libros[i].CopiasDisponibles--
			return true
		}
	}
	return false
}

// ManejadorUsuarios maneja la lógica relacionada con los usuarios
type ManejadorUsuarios struct {
	Usuarios []Usuario
}

func (m *ManejadorUsuarios) RegistrarUsuario(nombre, id, correo string) {
	usuario := Usuario{Nombre: nombre, NumeroIdentificacion: id, Correo: correo}
	m.Usuarios = append(m.Usuarios, usuario)
	fmt.Printf("Usuario registrado: %s\n", nombre)
}

func (m *ManejadorUsuarios) ObtenerUsuario(id string) (Usuario, bool) {
	for _, usuario := range m.Usuarios {
		if usuario.NumeroIdentificacion == id {
			return usuario, true
		}
	}
	return Usuario{}, false
}

// ManejadorPrestamos maneja la lógica relacionada con los préstamos
type ManejadorPrestamos struct {
	Prestamos []Prestamo
}

func (m *ManejadorPrestamos) ProcesarPrestamo(usuario Usuario, libro Libro) {
	prestamo := Prestamo{
		Usuario:       usuario,
		Libro:         libro,
		FechaPrestamo: time.Now(),
	}
	m.Prestamos = append(m.Prestamos, prestamo)
	fmt.Printf("Libro prestado: %s a %s\n", libro.Titulo, usuario.Nombre)
}

func practica26() {
	manejadorLibros := ManejadorLibros{}
	manejadorUsuarios := ManejadorUsuarios{}
	manejadorPrestamos := ManejadorPrestamos{}

	manejadorLibros.RegistrarLibro("El Quijote", "Miguel de Cervantes", 3)
	manejadorUsuarios.RegistrarUsuario("Juan Perez", "1234", "juan@example.com")

	libroDisponible := manejadorLibros.DisminuirCopia("El Quijote")
	if libroDisponible {
		usuario, encontrado := manejadorUsuarios.ObtenerUsuario("1234")
		if encontrado {
			manejadorPrestamos.ProcesarPrestamo(usuario, Libro{Titulo: "El Quijote"})
		}
	}
}
