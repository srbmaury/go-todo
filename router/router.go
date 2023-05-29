package router

import (
	todosService "Demo/internal/todos/service"
	userService "Demo/internal/users/service"
	"Demo/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	v2 := r.Group("/users")
	{
		// user routes
		v2.POST("/register", func(c *gin.Context) {
			userService.Register(c, db)
		})

		v2.POST("/login", func(c *gin.Context) {
			userService.Login(c, db)
		})
	}

	v1 := r.Group("/todos").Use(middlewares.JwtAuthMiddleware())
	{
		// todo routes
		v1.POST("/", func(c *gin.Context) {
			todosService.PostTodo(c, db)
		})
		v1.GET("", func(c *gin.Context) {
			todosService.GetTodos(c, db)
		})
		v1.GET("/:id", func(c *gin.Context) {
			todosService.GetTodo(c, db)
		})
		v1.PUT("/:id", func(c *gin.Context) {
			todosService.UpdateTodo(c, db)
		})
		v1.PATCH("/:id", func(c *gin.Context) {
			todosService.PartiallyUpdateTodo(c, db)
		})
		v1.DELETE("/:id", func(c *gin.Context) {
			todosService.DeleteTodo(c, db)
		})
	}
	return r
}
