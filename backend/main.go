package main

import (
	router "backend/app"
	"backend/clients"
	"backend/controllers"
	"backend/services"
	"log"

	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	// Inicializar la base de datos
	clients.InitDB()

	// Configurar el router
	r := router.SetupRouter()

	jwtService := services.NewJWTService()
    authController := controllers.NewAuthController(jwtService)

    r.POST("/register", authController.Register)
    r.POST("/login", authController.Login)

	protectedRoutes := r.Group("/api")
    protectedRoutes.Use(middleware.AuthorizeJWT(jwtService))
    protectedRoutes.GET("/protected", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "You are authorized"})
    })


	// Iniciar el servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
