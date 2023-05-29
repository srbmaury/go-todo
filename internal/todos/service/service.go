package todosService

import (
	todoModels "Demo/internal/todos/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostTodo(c *gin.Context, db *gorm.DB) {
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

	if err := db.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created"})
}

func GetTodos(c *gin.Context, db *gorm.DB) {
	var todos []todoModels.Todo

	if err := db.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context, db *gorm.DB) {
	todoID := c.Param("id")

	var todo todoModels.Todo
	if err := db.First(&todo, todoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context, db *gorm.DB) {
	todoID := c.Param("id")

	var todo todoModels.Todo
	if err := db.First(&todo, todoID).Error; err != nil {
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

func PartiallyUpdateTodo(c *gin.Context, db *gorm.DB) {
	todoID := c.Param("id")

	var todo todoModels.Todo
	if err := db.First(&todo, todoID).Error; err != nil {
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

func DeleteTodo(c *gin.Context, db *gorm.DB) {
	todoID := c.Param("id")

	var todo todoModels.Todo
	if err := db.First(&todo, todoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
