package main

import (
	"testing"
)

// TestCamposExistentes verifica que todos los campos existen
func TestCamposExistentes(t *testing.T) {
	datos := NewDatosPersonales("Tu Nombre", 30, "01/01/1990", []string{"Go", "Python", "Java"})

	if datos.Name == "" {
		t.Error("El campo 'Name' no existe o está vacío")
	}
	if datos.Age == 0 {
		t.Error("El campo 'Age' no existe o está vacío")
	}
	if datos.BirthDate == "" {
		t.Error("El campo 'BirthDate' no existe o está vacío")
	}
	if datos.ProgrammingLanguages == nil {
		t.Error("El campo 'ProgrammingLanguages' no existe o está vacío")
	}
}

// TestDatosCorrectos verifica que los datos introducidos son correctos
func TestDatosCorrectos(t *testing.T) {
	datos := NewDatosPersonales("Tu Nombre", 30, "01/01/1990", []string{"Go", "Python", "Java"})

	if datos.Name != "Tu Nombre" {
		t.Errorf("El campo 'Name' tiene un valor incorrecto: %s", datos.Name)
	}
	if datos.Age != 30 {
		t.Errorf("El campo 'Age' tiene un valor incorrecto: %d", datos.Age)
	}
	if datos.BirthDate != "01/01/1990" {
		t.Errorf("El campo 'BirthDate' tiene un valor incorrecto: %s", datos.BirthDate)
	}
	expectedLanguages := []string{"Go", "Python", "Java"}
	for i, lang := range datos.ProgrammingLanguages {
		if lang != expectedLanguages[i] {
			t.Errorf("El campo 'ProgrammingLanguages' tiene un valor incorrecto: %v", datos.ProgrammingLanguages)
		}
	}
}


