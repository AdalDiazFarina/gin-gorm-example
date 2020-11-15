package model

import (
	"time"
)

type Alquiler struct {
	ID        int       `gorm:"primaryKey"`
	DNI       string    `gorm:"type:varchar(50)"`
	Matricula string    `gorm:"matricula;type:varchar(10)"`
	Inicio    time.Time `gorm:"inicio"`
	Final     time.Time `gorm:"final"`
	Precio    float64   `gorm:"type:decimal(10,2)"`
	Coche     Coche     `gorm:"foreignKey:Matricula;references:Matricula"`
}

func (a *Alquiler) TableName() string {
	return "rent_a_car"
}
