package main

import (
	"github.com/FernandoMay/goempleados/controllers"
	"github.com/FernandoMay/goempleados/models"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/empleados", controllers.FindEmpleados)
	r.GET("/empleados/:id", controllers.FindEmpleado)
	r.POST("/empleados", controllers.CreateEmpleado)
	r.PATCH("/empleados/:id", controllers.UpdateEmpleado)
	r.DELETE("/empleados/:id", controllers.DeleteEmpleado)

	// Run the server
	r.Run()
}
