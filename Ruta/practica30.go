//* Solid: Principio de Inversión de Dependencias (DIP)

package main

import "fmt"

// Interfaz Notification
// Esta interfaz define el método Send que todas las implementaciones deben cumplir
type Notification interface {
	Send(message string)
}

// Implementación de Email
// Esta estructura implementa el método Send de la interfaz Notification
type EmailNotification struct{}

// Send implementa el envío de un mensaje por Email
func (e *EmailNotification) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// Implementación de PUSH
// Esta estructura implementa el método Send de la interfaz Notification
type PushNotification struct{}

// Send implementa el envío de un mensaje por PUSH
func (p *PushNotification) Send(message string) {
	fmt.Println("Sending PUSH Notification:", message)
}

// Implementación de SMS
// Esta estructura implementa el método Send de la interfaz Notification
type SMSNotification struct{}

// Send implementa el envío de un mensaje por SMS
func (s *SMSNotification) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// Sistema de Notificaciones
// Esta estructura usa la interfaz Notification para enviar notificaciones
type NotificationSystem struct {
	notifier Notification // La interfaz Notification que será utilizada
}

// NewNotificationSystem crea una nueva instancia de NotificationSystem
// con la implementación específica de Notification pasada como argumento
func NewNotificationSystem(notifier Notification) *NotificationSystem {
	return &NotificationSystem{notifier: notifier}
}

// Notify usa el método Send de la interfaz Notification para enviar el mensaje
func (ns *NotificationSystem) Notify(message string) {
	ns.notifier.Send(message)
}

func practica30() {
	// Crear instancias de cada implementación de Notification
	emailNotifier := &EmailNotification{}
	pushNotifier := &PushNotification{}
	smsNotifier := &SMSNotification{}

	// Crear sistemas de notificaciones con diferentes notificaciones
	// Aquí, el sistema de notificaciones no conoce la implementación específica,
	// solo sabe que está usando algo que cumple con la interfaz Notification
	emailNotificationSystem := NewNotificationSystem(emailNotifier)
	pushNotificationSystem := NewNotificationSystem(pushNotifier)
	smsNotificationSystem := NewNotificationSystem(smsNotifier)

	// Enviar notificaciones usando diferentes sistemas de notificaciones
	emailNotificationSystem.Notify("Hello via Email!")
	pushNotificationSystem.Notify("Hello via PUSH!")
	smsNotificationSystem.Notify("Hello via SMS!")
}
