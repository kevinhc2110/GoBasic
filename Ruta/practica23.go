//* Patrones de diseño: Singleton

package main

import (
	"fmt"
	"sync"
)

/*
DIFICULTAD EXTRA:
Utiliza el patrón de diseño "singleton" para representar una clase que
haga referencia a la sesión de usuario de una aplicación ficticia.
La sesión debe permitir asignar un usuario (id, username, nombre y email),
recuperar los datos del usuario y borrar los datos de la sesión.
*/

// User representa los datos del usuario.
type User struct {
	ID       int
	Username string
	Nombre   string
	Email    string
	Activo   bool // Indica si la sesión está activa.
}

// UserSession es la estructura singleton que maneja la sesión del usuario.
type UserSession struct {
	user *User
	// mu es un mutex para asegurar que el acceso a la sesión del usuario sea seguro en entornos concurrentes.
	mu sync.Mutex // Mutex para manejar concurrencia en acceso a la sesión.
}

// Usa sync.Once para asegurar que UserSession se inicialice una sola vez.
var (
	instance *UserSession
	once     sync.Once
)

// GetInstance retorna la única instancia de UserSession.
func GetInstance() *UserSession {
	once.Do(func() {
		instance = &UserSession{}
	})
	return instance
}

// SetUser asigna un usuario a la sesión.
// Usa un mutex para asegurar que la operación sea segura en entornos concurrentes.
func (s *UserSession) SetUser(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user.Activo = true // Marcar la sesión como activa.
	s.user = user
}

// GetUser retorna los datos del usuario de la sesión.
func (s *UserSession) GetUser() *User {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.user != nil && s.user.Activo {
		return s.user
	}
	return nil
}

// ClearUser borra los datos de la sesión.
func (s *UserSession) ClearUser() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.user != nil {
		s.user.Activo = false // Marcar la sesión como inactiva.
		s.user = nil
	}
}

func practica23() {
	// Obtener la instancia de la sesión
	session := GetInstance()

	// Crear un nuevo usuario
	user := &User{
		ID:       1,
		Username: "johndoe",
		Nombre:   "John Doe",
		Email:    "john.doe@example.com",
	}

	// Asignar el usuario a la sesión
	session.SetUser(user)

	// Recuperar el usuario de la sesión
	retrievedUser := session.GetUser()
	if retrievedUser != nil {
		fmt.Printf("Usuario en sesión: %v\n", *retrievedUser)
	} else {
		fmt.Println("No hay usuario en sesión.")
	}

	// Borrar la sesión del usuario
	session.ClearUser()

	// Verificar que la sesión ha sido borrada
	retrievedUser = session.GetUser()
	if retrievedUser != nil {
		fmt.Printf("Usuario en sesión: %v\n", *retrievedUser)
	} else {
		fmt.Println("No hay usuario en sesión.")
	}
}
