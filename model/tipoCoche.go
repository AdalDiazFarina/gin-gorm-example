package model

type TipoCoche struct {
	IDTipo          int     `gorm:"column:idTipo;primaryKey"`
	DescripcionTipo string  `gorm:"column:DescripcionTipo;type:varchar(100)"`
	PrecioDia       float64 `gorm:"column:PrecioDia;type:decimal(10,2)"`
}

func (t *TipoCoche) TableName() string {
	return "tipoCoche"
}
