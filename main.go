package main

import (
	"log"
	"os"

	"golangtodolist/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)

	router.Run(":" + os.Getenv("PORT"))

}
