package main

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "wallet-sdk/internal/config"
    "wallet-sdk/internal/database"
    "wallet-sdk/pkg/users"
)

func main() {
    config.LoadConfig()
    database.Connect()

    // Initialize repository and service
    userRepo := users.NewRepository(database.Instance)
    userService := users.NewService(userRepo)

    r := gin.Default()

    r.POST("/users", func(c *gin.Context) {
        var user users.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        if err := userService.CreateUser(&user); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, user)
    })

    r.GET("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid user ID"})
            return
        }
        
        user, err := userService.GetUserByID(id)
        if err != nil {
            c.JSON(404, gin.H{"error": "User not found"})
            return
        }
        c.JSON(200, user)
    })

    r.Run(":8080")
}
