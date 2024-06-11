package main

import (
    "go-data/db"
    "go-data/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    db.Init()

    r := gin.Default()

    r.POST("/users", handlers.CreateUser)

    r.Run(":8080")
}
