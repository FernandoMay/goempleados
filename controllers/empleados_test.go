package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/FernandoMay/goempleados/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}
	db.AutoMigrate(&models.Empleado{})
	models.DB = db
	os.Exit(m.Run())
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.GET("/api/empleados", FindEmpleados)
	r.GET("/api/empleados/:id", FindEmpleado)
	r.POST("/api/empleados", CreateEmpleado)
	r.PATCH("/api/empleados/:id", UpdateEmpleado)
	r.DELETE("/api/empleados/:id", DeleteEmpleado)
	return r
}

func cleanup() {
	models.DB.Exec("DELETE FROM empleados")
}

func validEmpleadoJSON() string {
	return `{"nombre":"Juan","apellidop":"Perez","apellidom":"Garcia","area":"IT","fechanac":"1990-01-01T00:00:00Z","sueldo":50000}`
}

func TestCreateEmpleado(t *testing.T) {
	cleanup()
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/empleados", strings.NewReader(validEmpleadoJSON()))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp["data"].(map[string]interface{})
	if data["nombre"] != "Juan" {
		t.Errorf("expected nombre 'Juan', got %v", data["nombre"])
	}
	if data["apellidop"] != "Perez" {
		t.Errorf("expected apellidop 'Perez', got %v", data["apellidop"])
	}
	if data["area"] != "IT" {
		t.Errorf("expected area 'IT', got %v", data["area"])
	}
	if data["sueldo"] != 50000 {
		t.Errorf("expected sueldo 50000, got %v", data["sueldo"])
	}
}

func TestCreateEmpleado_Invalid(t *testing.T) {
	cleanup()
	r := setupRouter()
	w := httptest.NewRecorder()
	body := `{"nombre":""}`
	req, _ := http.NewRequest("POST", "/api/empleados", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestFindEmpleados_Empty(t *testing.T) {
	cleanup()
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/empleados", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp["data"].([]interface{})
	if len(data) != 0 {
		t.Errorf("expected empty list, got %d items", len(data))
	}
}

func TestFindEmpleados(t *testing.T) {
	cleanup()
	now := time.Now()
	models.DB.Create(&models.Empleado{Nombre: "Ana", ApellidoP: "Lopez", ApellidoM: "Ruiz", Area: "HR", FechaNac: now, Sueldo: 40000})
	models.DB.Create(&models.Empleado{Nombre: "Luis", ApellidoP: "Martinez", ApellidoM: "Cruz", Area: "IT", FechaNac: now, Sueldo: 60000})

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/empleados", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp["data"].([]interface{})
	if len(data) != 2 {
		t.Errorf("expected 2 empleados, got %d", len(data))
	}
}

func TestFindEmpleado(t *testing.T) {
	cleanup()
	emp := models.Empleado{Nombre: "Carlos", ApellidoP: "Diaz", ApellidoM: "Soto", Area: "Finance", FechaNac: time.Now(), Sueldo: 55000}
	models.DB.Create(&emp)

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/empleados/%d", emp.ID), nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp["data"].(map[string]interface{})
	if data["nombre"] != "Carlos" {
		t.Errorf("expected nombre 'Carlos', got %v", data["nombre"])
	}
}

func TestFindEmpleado_NotFound(t *testing.T) {
	cleanup()
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/empleados/999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestUpdateEmpleado(t *testing.T) {
	cleanup()
	emp := models.Empleado{Nombre: "Original", ApellidoP: "Name", ApellidoM: "X", Area: "Sales", FechaNac: time.Now(), Sueldo: 30000}
	models.DB.Create(&emp)

	r := setupRouter()
	w := httptest.NewRecorder()
	body := `{"nombre":"Updated Name","sueldo":45000}`
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/empleados/%d", emp.ID), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp["data"].(map[string]interface{})
	if data["nombre"] != "Updated Name" {
		t.Errorf("expected nombre 'Updated Name', got %v", data["nombre"])
	}
}

func TestUpdateEmpleado_NotFound(t *testing.T) {
	cleanup()
	r := setupRouter()
	w := httptest.NewRecorder()
	body := `{"nombre":"Nope"}`
	req, _ := http.NewRequest("PATCH", "/api/empleados/999", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestDeleteEmpleado(t *testing.T) {
	cleanup()
	emp := models.Empleado{Nombre: "Delete", ApellidoP: "Me", ApellidoM: "Now", Area: "Ops", FechaNac: time.Now(), Sueldo: 35000}
	models.DB.Create(&emp)

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/empleados/%d", emp.ID), nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["data"] != true {
		t.Errorf("expected true, got %v", resp["data"])
	}

	count := int64(0)
	models.DB.Model(&models.Empleado{}).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 after delete, got %d", count)
	}
}

func TestDeleteEmpleado_NotFound(t *testing.T) {
	cleanup()
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/empleados/999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}
