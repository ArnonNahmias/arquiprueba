package domain

type Course struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Nombre     string `json:"nombre"`
	Dificultad string `json:"dificultad"`
	Precio     int    `json:"precio"`
	Direccion  string `json:"direccion"`
}
