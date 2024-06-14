package main

import (
	"backend/app"
	"backend/clients"
	"log"
)

func main() {
	// Inicializar la base de datos
	clients.InitDB()

	// Configurar el router
	r := router.SetupRouter()

	// Iniciar el servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
