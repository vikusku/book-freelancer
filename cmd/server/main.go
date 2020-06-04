package main

import (
	"github.com/gin-gonic/gin"

	"github.com/vikusku/book-freelancer/internal/handlers"
	"log"
)

var PORT = ":7777"

func main() {
	router := gin.Default()
	router.GET("/ping", handlers.Ping)

	err := router.Run(PORT)
	if err != nil {
		log.Fatalln(err)
	}
}