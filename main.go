package main

import (
	"fmt"
	_ "net/http"

	"github.com/cheildo/jwt-auth-golang/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", handlers.LoginHandler)
	router.GET("/protected", handlers.ProtectedHandler)

	fmt.Println("Starting the server")
	err := router.Run("localhost:4000")
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
}
