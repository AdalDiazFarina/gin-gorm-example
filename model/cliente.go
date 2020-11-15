package model

import (
	"time"
)

type Cliente struct {
	DNI              string     `gorm:"type:varchar(9);primaryKey"`
	Apellidos        string     `gorm:"apellidos;type:varchar(50)"`
	Nombre           string     `gorm:"nombre;type:varchar(50)"`
	Datos            string     `gorm:"datos;type:varchar(50)"`
	Nacimiento       time.Time  `gorm:"nacimiento"`
	ExpedicionCarnet time.Time  `gorm:"expedicionCarnet;"`
	Alquileres       []Alquiler `gorm:"foreignKey:DNI"`
}

func (a *Cliente) TableName() string {
	return "cliente"
}
