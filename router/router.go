package router

import (
	todosService "Demo/internal/todos/service"
	userService "Demo/internal/users/service"
	middlewares "Demo/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	v2 := r.Group("/users")
	{
		// user routes
		v2.POST("/register", userService.Register(db))
		v2.POST("/login", userService.Login(db))
	}

	v1 := r.Group("/todos").Use(middlewares.JwtAuthMiddleware())
	{
		// todo routes
		v1.POST("/", todosService.PostTodo(db))
		v1.GET("", todosService.GetTodos(db))
		v1.GET("/:id", todosService.GetTodo(db))
		v1.PUT("/:id", todosService.UpdateTodo(db))
		v1.PATCH("/:id", todosService.PartiallyUpdateTodo(db))
		v1.DELETE("/:id", todosService.DeleteTodo(db))
	}

	return r
}
