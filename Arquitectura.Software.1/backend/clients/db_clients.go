package clients

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:admin@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	log.Println("Database connected successfully")

	Migrate()
	SeedDB()
}

func Migrate() {
	log.Println("Migrating database...")
	err := DB.AutoMigrate(&Course{}, &Usuario{}, &Subscription{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully")
}

func SeedDB() {
	log.Println("Seeding database...")
	cursos := []Course{
		{Nombre: "Ingles B2", Dificultad: "Medio", Precio: 45, Direccion: "José Roque Funes 1511 X5000ABE Córdoba"},
		{Nombre: "Hacking Etico", Dificultad: "Dificil", Precio: 60, Direccion: "Paseo de la Reforma 505, CDMX"},
	}
	for _, curso := range cursos {
		DB.Create(&curso)
	}
	log.Println("Database seeded successfully")
}

type Course struct {
	ID         uint `gorm:"primaryKey"`
	Nombre     string
	Dificultad string
	Precio     int
	Direccion  string
}

type Usuario struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

type Subscription struct {
	ID       uint `gorm:"primaryKey"`
	UserID   int
	CourseID int
}
