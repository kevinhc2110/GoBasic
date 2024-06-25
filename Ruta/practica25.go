package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
DIFICULTAD EXTRA:
Crea un programa ficticio de gestión de tareas que permita añadir, eliminar
y listar dichas tareas.
- Añadir: recibe nombre y descripción.
- Eliminar: por nombre de la tarea.
Implementa diferentes mensajes de log que muestren información según la
tarea ejecutada (a tu elección).
Utiliza el log para visualizar el tiempo de ejecución de cada tarea.
*/

// Estructura para representar una tarea

type Task struct {
	Name        string
	Description string
}

// Slice para almacenar las tareas

var tasks []Task

// Función de inicialización para configurar logrus
func init() {
	// Configuración del formateador de log para incluir marcas de tiempo completas
	log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
	})
	// Establecer el nivel de log en DebugLevel para capturar todos los niveles de logs
	log.SetLevel(log.DebugLevel)
}

func addTask(name string, description string) {
	// Registrar el tiempo de inicio de la operación
	start := time.Now()
	// Crear una nueva tarea y añadirla al slice de tareas
	task := Task{Name: name, Description: description}
	tasks = append(tasks, task)
	 // Registrar mensajes de información y depuración
	log.Infof("Tarea añadida: %s", name)
	log.Debugf("Descripción: %s", description)
	// Calcular y registrar el tiempo de ejecución de la operación
	elapsed := time.Since(start)
	log.Infof("Tiempo de ejecución de añadir tarea: %s", elapsed)
}

func removeTask(name string) {
	start := time.Now()
	// Buscar la tarea en el slice
	index := -1
	for i, task := range tasks {
			if task.Name == name {
					index = i
					break
			}
	}
	 // Si se encuentra la tarea, eliminarla del slice
	if index != -1 {
			tasks = append(tasks[:index], tasks[index+1:]...)
			log.Infof("Tarea eliminada: %s", name)
	} else {
		 // Si no se encuentra, registrar una advertencia
			log.Warnf("Tarea no encontrada: %s", name)
	}
	elapsed := time.Since(start)
	log.Infof("Tiempo de ejecución de eliminar tarea: %s", elapsed)
}

func listTasks() {
	start := time.Now()
	if len(tasks) == 0 {
			log.Warn("No hay tareas para mostrar")
	} else {
			log.Info("Listando tareas:")
			for _, task := range tasks {
					fmt.Printf("Nombre: %s, Descripción: %s\n", task.Name, task.Description)
			}
	}
	elapsed := time.Since(start)
	log.Infof("Tiempo de ejecución de listar tareas: %s", elapsed)
}

func practica25() {
	addTask("Comprar leche", "Comprar leche en el supermercado")
	addTask("Estudiar Go", "Leer el libro 'The Go Programming Language'")
	listTasks()
	removeTask("Comprar leche")
	listTasks()
	removeTask("Ir al gimnasio")
}