package handlers

import (
	"fmt"
	"net/http"

	"github.com/cheildo/jwt-auth-golang/auth"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var u User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fmt.Printf("The user request value %v", u)

	if u.Username == "talal" && u.Password == "123456" {
		tokenString, err := auth.CreateToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid json format"})
}

func ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "protected area"})
}
