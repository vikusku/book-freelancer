package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vikusku/book-freelancer/internal/handler"
	"log"
)

var host, port = "localhost", "7777"

func Run() {
	router := gin.Default()
	router.GET("/ping", handler.Ping)

	endpoint := host + ":" + port

	log.Printf("connect to http://%s/ for GraphQL playground", endpoint)
	router.GET("/", handler.Playground("/query"))
	router.POST("/query", handler.QueryHandler())

	err := router.Run(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
}
