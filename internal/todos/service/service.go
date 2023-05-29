package todosService

import (
	contextapi "Todos/context"
	todoModels "Todos/internal/todos/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostTodo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var uid contextapi.MyContextKey
		userId := contextapi.GetValue(uid)

		var newTodo todoModels.Todo

		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.AutoMigrate(&todoModels.Todo{}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to migrate Todo table"})
			return
		}
		newTodo.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
		newTodo.UpdatedAt = time.Now().Format("01-02-2006 15:04:05")
		newTodo.UserId = userId

		if err := db.Create(&newTodo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Todo created"})
	}
}

func GetTodos(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var uid contextapi.MyContextKey
		userId := contextapi.GetValue(uid)

		var todos []todoModels.Todo

		if err := db.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
			return
		}

		c.JSON(http.StatusOK, todos)
	}
}

func GetTodo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var uid contextapi.MyContextKey
		userId := contextapi.GetValue(uid)

		todoID := c.Param("id")

		var todo todoModels.Todo
		if err := db.Where("user_id = ?", userId).First(&todo, todoID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var uid contextapi.MyContextKey
		userId := contextapi.GetValue(uid)

		todoID := c.Param("id")

		var todo todoModels.Todo
		if err := db.Where("user_id = ?", userId).First(&todo, todoID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		var updatedTodo todoModels.Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo.Title = updatedTodo.Title
		todo.Description = updatedTodo.Description
		todo.UpdatedAt = time.Now().Format("01-02-2006 15:04:05")

		if err := db.Save(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func PartiallyUpdateTodo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var uid contextapi.MyContextKey
		userId := contextapi.GetValue(uid)
		todoID := c.Param("id")

		var todo todoModels.Todo
		if err := db.Where("user_id = ?", userId).First(&todo, todoID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		var updatedData map[string]interface{}

		if err := c.ShouldBindJSON(&updatedData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedData["UpdatedAt"] = time.Now().Format("01-02-2006 15:04:05")
		if err := db.Model(&todo).Updates(updatedData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var uid contextapi.MyContextKey
		userId := contextapi.GetValue(uid)
		todoID := c.Param("id")

		var todo todoModels.Todo
		if err := db.Where("user_id = ?", userId).First(&todo, todoID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		if err := db.Delete(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
	}
}
