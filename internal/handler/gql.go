package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vikusku/book-freelancer/internal/graph"
	"github.com/vikusku/book-freelancer/internal/graph/generated"
	"log"
)

func QueryHandler(dbConnection *gorm.DB) gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{
				DbConnection: dbConnection,
			},
		}))

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))

		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "referer,content-type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func Playground(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", path)

	return func(c *gin.Context) {
		log.Println("playground handler")
		h.ServeHTTP(c.Writer, c.Request)
	}
}