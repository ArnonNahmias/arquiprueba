package app

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Grupo de rutas para cursos
	courseGroup := r.Group("/courses")
	{
		courseGroup.GET("/", controllers.GetCourses)
		courseGroup.GET("/:id", controllers.GetCourseByID)
		courseGroup.POST("/", controllers.CreateCourse)
		courseGroup.PUT("/:id", controllers.UpdateCourse)
		courseGroup.DELETE("/:id", controllers.DeleteCourse)
	}

	// Rutas adicionales según sea necesario
	r.GET("/search", controllers.Search)
	r.POST("/subscribe", controllers.Subscribe)

	return r
}
