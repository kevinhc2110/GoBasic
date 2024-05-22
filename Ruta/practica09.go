package main

import "fmt"

/*
DIFICULTAD EXTRA:
Implementa la jerarquía de una empresa de desarrollo formada por Empleados que
pueden ser Gerentes, Gerentes de Proyectos o Programadores.
Cada empleado tiene un identificador y un nombre.
Dependiendo de su labor, tienen propiedades y funciones exclusivas de su
actividad, y almacenan los empleados a su cargo.
*/

// Definir la interfaz Empleado

type Empleado interface {
	Identificador() int
	Nombre() string
	MostrarInformacion()
	AgregarSubordinado(Empleado)
	ObtenerSubordinados() []Empleado
}

// Estructura base para Empleado

type EmpleadoBase struct {
	id            int
	name          string
	subordinados  []Empleado
}

// Métodos comunes para todos los empleados

func (e *EmpleadoBase) Identificador() int {
	return e.id
}

func (e *EmpleadoBase) Nombre() string {
	return e.name
}

func (e *EmpleadoBase) AgregarSubordinado(subordinado Empleado) {
	e.subordinados = append(e.subordinados, subordinado)
}

func (e *EmpleadoBase) ObtenerSubordinados() []Empleado {
	return e.subordinados
}

// Estructura para Gerente

type Gerente struct {
	EmpleadoBase
	Departamento string
}

// Método para mostrar información específica de Gerente

func (g Gerente) MostrarInformacion() {
	fmt.Printf("Gerente ID: %d, Nombre: %s, Departamento: %s\n", g.id, g.name, g.Departamento)
	for _, subordinado := range g.subordinados {
		fmt.Printf("  Subordinado -> ")
		subordinado.MostrarInformacion()
	}
}

// Métodos específicos para Gerente

func (g *Gerente) AsignarProyecto(subordinado *GerenteProyectos, proyecto string) {
	fmt.Printf("Asignando proyecto %s a %s...\n", proyecto, subordinado.Nombre())
	// Lógica para asignar proyecto
}

// Estructura para Gerente de Proyectos

type GerenteProyectos struct {
	EmpleadoBase
	Proyecto string
}

// Método para mostrar información específica de Gerente de Proyectos

func (gp GerenteProyectos) MostrarInformacion() {
	fmt.Printf("Gerente de Proyectos ID: %d, Nombre: %s, Proyecto: %s\n", gp.id, gp.name, gp.Proyecto)
	for _, subordinado := range gp.subordinados {
		fmt.Printf("  Subordinado -> ")
		subordinado.MostrarInformacion()
	}
}

// Métodos específicos para Gerente de Proyecto

func (gp *GerenteProyectos) PlanificarSprint() {
	fmt.Printf("Planificando sprint para el proyecto %s...\n", gp.Proyecto)
	// Lógica para planificar sprint
}

// Estructura para Programador

type Programador struct {
	EmpleadoBase
	Lenguaje string
}

// Método para mostrar información específica de Programador

func (p Programador) MostrarInformacion() {
	fmt.Printf("Programador ID: %d, Nombre: %s, Lenguaje: %s\n", p.id, p.name, p.Lenguaje)
	for _, subordinado := range p.subordinados {
		fmt.Printf("  Subordinado -> ")
		subordinado.MostrarInformacion()
	}
}

// Métodos específicos para Programador

func (p *Programador) EscribirCodigo() {
	fmt.Printf("%s está escribiendo código en %s...\n", p.Nombre(), p.Lenguaje)
	// Lógica para escribir código
}

func practica09()/*main*/ {

	// Crear instancias de Gerente, Gerente de Proyectos y Programador

	gerente := Gerente{EmpleadoBase: EmpleadoBase{id: 1, name: "Carlos"}, Departamento: "Ventas"}
	gerenteProyectos := GerenteProyectos{EmpleadoBase: EmpleadoBase{id: 2, name: "Ana"}, Proyecto: "Proyecto A"}
	programador1 := Programador{EmpleadoBase: EmpleadoBase{id: 3, name: "Juan"}, Lenguaje: "Go"}
	programador2 := Programador{EmpleadoBase: EmpleadoBase{id: 4, name: "María"}, Lenguaje: "Python"}

	// Agregar subordinados

	gerente.AgregarSubordinado(&gerenteProyectos)
	gerenteProyectos.AgregarSubordinado(&programador1)
	gerenteProyectos.AgregarSubordinado(&programador2)

	// Crear una lista de empleados

	empleados := []Empleado{&gerente, &gerenteProyectos, &programador1, &programador2}

	// Mostrar la información de todos los empleados

	for _, empleado := range empleados {
		empleado.MostrarInformacion()
	}

	// Ejemplo de uso de los métodos específicos
	
	gerente.AsignarProyecto(&gerenteProyectos, "Nuevo Proyecto")
	gerenteProyectos.PlanificarSprint()
	programador1.EscribirCodigo()
}
