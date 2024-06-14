package dao

import (
	"time"
)

type Course struct {
	IdCurso    int       `gorm:"primaryKey;column:Id_curso;autoIncrement"`
	Nombre     string    `gorm:"column:Nombre;not null"`
	Dificultad string    `gorm:"column:Dificultad;not null"`
	Precio     int       `gorm:"column:Precio;not null"`
	Direccion  string    `gorm:"column:Direccion"`
	ImageURL   string    `gorm:"column:ImageURL"`
	CreatedAt  time.Time `gorm:"column:Created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:Updated_at;autoUpdateTime"`
}
