package userService

import (
	userModels "Demo/internal/users/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var user userModels.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u := userModels.User{}
		u.Username = user.Username
		u.Password = user.Password
		u.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
		u.UpdatedAt = time.Now().Format("01-02-2006 15:04:05")

		user.BeforeSave()
		_, err := user.SaveUser(db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "registration success"})
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input LoginInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u := userModels.User{}

		u.Username = input.Username
		u.Password = input.Password

		token, err := userModels.LoginCheck(u.Username, u.Password, db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
