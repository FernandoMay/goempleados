package controllers

import (
	"net/http"
	"time"

	"github.com/FernandoMay/goempleados/models"
	"github.com/gin-gonic/gin"
)

type CreateEmpleadoInput struct {
	Nombre    string    `json:"nombre" binding:"required"`
	ApellidoP string    `json:"apellidop" binding:"required"`
	ApellidoM string    `json:"apellidom" binding:"required"`
	Area      string    `json:"area" binding:"required"`
	FechaNac  time.Time `json:"fechanac" binding:"required"`
	Sueldo    uint      `json:"sueldo" binding:"required"`
}

type UpdateEmpleadoInput struct {
	ID        uint      `json:"id"`
	Nombre    string    `json:"nombre"`
	ApellidoP string    `json:"apellidop"`
	ApellidoM string    `json:"apellidom"`
	Area      string    `json:"area"`
	FechaNac  time.Time `json:"fechanac"`
	Sueldo    uint      `json:"sueldo"`
}

// GET /Empleados
// Find all Empleados
func FindEmpleados(c *gin.Context) {
	var Empleados []models.Empleado
	models.DB.Find(&Empleados)

	c.JSON(http.StatusOK, gin.H{"data": Empleados})
}

// GET /Empleados/:id
// Find a Empleado
func FindEmpleado(c *gin.Context) {
	// Get model if exist
	var Empleado models.Empleado
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Empleado).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Empleado})
}

// POST /Empleados
// Create new Empleado
func CreateEmpleado(c *gin.Context) {
	// Validate input
	var input CreateEmpleadoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Empleado
	Empleado := models.Empleado{Nombre: input.Nombre, ApellidoP: input.ApellidoP, ApellidoM: input.ApellidoM, Area: input.Area, FechaNac: input.FechaNac, Sueldo: uint(input.Sueldo)}
	models.DB.Create(&Empleado)

	c.JSON(http.StatusOK, gin.H{"data": Empleado})
}

// PATCH /Empleados/:id
// Update a Empleado
func UpdateEmpleado(c *gin.Context) {
	// Get model if exist
	var Empleado models.Empleado
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Empleado).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateEmpleadoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Empleado).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": Empleado})
}

// DELETE /Empleados/:id
// Delete a Empleado
func DeleteEmpleado(c *gin.Context) {
	// Get model if exist
	var Empleado models.Empleado
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Empleado).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Empleado)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
