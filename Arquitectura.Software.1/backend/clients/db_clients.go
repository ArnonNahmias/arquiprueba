package clients

import (
	"backend/dao"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:pass@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb4&parseTime=True&loc=Local"
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
	DB.AutoMigrate(&dao.Usuario{}, &dao.Course{}, &dao.Subscription{})
}

func SeedDB() {
	log.Println("Seeding database...")
	admin := dao.Usuario{NombreUsuario: "admin", Contrasena: "admin", Tipo: "admin"}
	user := dao.Usuario{NombreUsuario: "user", Contrasena: "user", Tipo: "normal"}

	DB.FirstOrCreate(&admin, dao.Usuario{NombreUsuario: "	admin"})
	DB.FirstOrCreate(&user, dao.Usuario{NombreUsuario: "user"})
}
