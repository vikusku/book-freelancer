package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vikusku/book-freelancer/internal/handler"
	"github.com/vikusku/book-freelancer/internal/orm"
	"log"
)

func Run() {
	dbConnection := orm.OpenDB()

	router := gin.Default()
	router.GET("/ping", handler.Ping)


	log.Printf("connect to for GraphQL playground")
	router.GET("/", handler.Playground("/query"))
	router.POST("/query", handler.QueryHandler(dbConnection))
	router.OPTIONS("/query", handler.QueryHandler(dbConnection))

	err := router.Run(":7777")
	if err != nil {
		log.Fatalln(err)
	}
}
