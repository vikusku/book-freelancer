package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vikusku/book-freelancer/internal/handler"
	"github.com/vikusku/book-freelancer/internal/orm"
	"log"
)

var host, port = "localhost", "7777"

func Run() {
	dbConnection := orm.OpenDB()

	router := gin.Default()
	router.GET("/ping", handler.Ping)

	endpoint := host + ":" + port

	log.Printf("connect to http://%s/ for GraphQL playground", endpoint)
	router.GET("/", handler.Playground("/query"))
	router.POST("/query", handler.QueryHandler(dbConnection))
	router.OPTIONS("/query", handler.QueryHandler(dbConnection))

	err := router.Run(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
}
