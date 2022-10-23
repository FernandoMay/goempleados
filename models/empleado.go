package models

import "time"

type Empleado struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Nombre    string    `json:"nombre"`
	ApellidoP string    `json:"apellidop"`
	ApellidoM string    `json:"apellidom"`
	Area      string    `json:"area"`
	FechaNac  time.Time `json:"fechanac"`
	Sueldo    float32   `json:"sueldo"`
}
