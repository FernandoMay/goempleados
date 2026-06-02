package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEmpleadoStruct(t *testing.T) {
	now := time.Now()
	emp := Empleado{
		Nombre:    "Juan",
		ApellidoP: "Perez",
		ApellidoM: "Garcia",
		Area:      "IT",
		FechaNac:  now,
		Sueldo:    50000,
	}

	if emp.Nombre != "Juan" {
		t.Errorf("expected Nombre 'Juan', got '%s'", emp.Nombre)
	}
	if emp.ApellidoP != "Perez" {
		t.Errorf("expected ApellidoP 'Perez', got '%s'", emp.ApellidoP)
	}
	if emp.ApellidoM != "Garcia" {
		t.Errorf("expected ApellidoM 'Garcia', got '%s'", emp.ApellidoM)
	}
	if emp.Area != "IT" {
		t.Errorf("expected Area 'IT', got '%s'", emp.Area)
	}
	if !emp.FechaNac.Equal(now) {
		t.Errorf("expected FechaNac %v, got %v", now, emp.FechaNac)
	}
	if emp.Sueldo != 50000 {
		t.Errorf("expected Sueldo 50000, got %d", emp.Sueldo)
	}
	if emp.ID != 0 {
		t.Errorf("expected ID 0, got %d", emp.ID)
	}
}

func TestEmpleadoJSON(t *testing.T) {
	now := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	emp := Empleado{ID: 1, Nombre: "Maria", ApellidoP: "Lopez", ApellidoM: "Ramos", Area: "HR", FechaNac: now, Sueldo: 40000}

	data, err := json.Marshal(emp)
	if err != nil {
		t.Fatal(err)
	}

	var decoded Empleado
	json.Unmarshal(data, &decoded)
	if decoded.Nombre != "Maria" {
		t.Errorf("expected Nombre 'Maria', got '%s'", decoded.Nombre)
	}
	if decoded.Sueldo != 40000 {
		t.Errorf("expected Sueldo 40000, got %d", decoded.Sueldo)
	}
	if decoded.ID != 1 {
		t.Errorf("expected ID 1, got %d", decoded.ID)
	}
}
