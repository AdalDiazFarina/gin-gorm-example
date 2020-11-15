package model

import (
	"log"

	"github.com/manolors/full-api-example/storage"
)

type Coche struct {
	Matricula string `gorm:"column:matricula;type:varchar(10);primaryKey"`
	Estado    string `gorm:"column:estado;type:varchar(100)"`
	IDTipo    int    `gorm:"column:idTipo;-"`
	TipoCoche `gorm:"foreignKey:idTipo;references:idTipo"`
}

type Coches []Coche

func (c *Coches) Get() error {
	rows, err := storage.DB.Debug().
		Model(&Coche{}).
		Select(`coche.matricula,
						coche.estado,
						tipoCoche.DescripcionTipo,
						tipoCoche.PrecioDia`).
		Joins("JOIN tipoCoche on tipoCoche.IDTipo = coche.IDTipo").
		Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		coche := Coche{}
		err = storage.DB.ScanRows(rows, &coche)
		if err != nil {
			log.Println("error al bindear", err)
		} else {
			log.Printf("Coche:%v\n", coche)
		}
		*c = append([]Coche(*c), coche)
	}
	return nil

}

func (c *Coche) Nuevo() error {
	result := storage.DB.Debug().Save(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Coche) TableName() string {
	return "coche"
}
