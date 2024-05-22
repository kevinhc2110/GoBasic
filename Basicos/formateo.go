package main

import (
	"log"
	"os"
	"text/template"
)

func formateo() /*main*/ {

	println(
		"Enteros: %d",
		"Flotantes: %f",
		"Cadenas: %s",
		"Booleano: %t",
		"Valores predeterminados: %v",
		"Ancho para enteros: %3d",
		"Ancho y precisión para flotantes: %.2f",
		"Relleno a la izquierda: %3s",
		"Signo y relleno a la izquierda: %-3s",
		"Hexadecimal: %#x\n", 0xFF,
		"Separador de miles: %d\n", 123456789,
		"Caracter unicode: %c",
		"Entero sin signo en base 16: %x o %X",
		"Entero sin signo en base 2: %b",
		"Entero sin signo en base 8: %o",
		"Cadena entre comillas dobles con caracteres escapados según el formato Go: %q",
		"Fechas y horas: %t, %T, %v",
		"Números complejos: %e, %E, %f, %F",
		"Puntero, imprime la representación en notación hexadecimal del puntero: %p",
		"Cadena Unicode: %U",
		"Tipo, imprime el tipo del valor: %T",
		" Número de punto flotante en formato de punto flotante o de exponente, dependiendo del valor: %g, %G",
	)

	// Plantillas de texto:

	nombre := "María González"
	edad := 28
	profesion := "Desarrolladora de software"

	tmpl := "Nombre: {{.Nombre}}\nEdad: {{.Edad}} años\nProfesión: {{.Profesion}}"

	data := map[string]interface{}{
		"Nombre":    nombre,
		"Edad":      edad,
		"Profesion": profesion,
	}

	t, err := template.New("persona").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

	// En este ejemplo, se define una plantilla de texto tmpl con marcadores de posición {{.Nombre}}, {{.Edad}} y {{.Profesion}}. Luego, se crea un mapa data con los valores correspondientes a cada marcador. Finalmente, se utiliza la función template.New() para crear una plantilla a partir de tmpl y la función Execute() para renderizar la plantilla con los datos del mapa data.
}