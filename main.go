package main

import (
	"github.com/FernandoMay/goempleados/controllers"
	"github.com/FernandoMay/goempleados/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/api/empleados", controllers.FindEmpleados)
	r.GET("/api/empleados/:id", controllers.FindEmpleado)
	r.POST("/api/empleados", controllers.CreateEmpleado)
	r.PATCH("/api/empleados/:id", controllers.UpdateEmpleado)
	r.DELETE("/api/empleados/:id", controllers.DeleteEmpleado)

	// Run the server
	r.Run()
}
