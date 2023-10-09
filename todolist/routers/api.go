package routers

import (
	_ "github.com/elysiamori/finalproject1/kelompok6/docs"

	"github.com/elysiamori/finalproject1/kelompok6/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartRouter
// @title Todolist APP
// @description Todos API Kelompok 6
// @github github.com/elysiamori
// @version 1
// @host localhost:3000
// @BasePath /api
func StartRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/todos", controllers.GetAllTodos)
	api.GET("/todos/:id", controllers.GetTodosID)
	api.POST("/todos", controllers.AddTodos)
	api.PUT("/todos/:id", controllers.UpdateTodos)
	api.DELETE("/todos/:id", controllers.DeleteTodos)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
