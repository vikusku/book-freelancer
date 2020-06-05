package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vikusku/book-freelancer/internal/gql/generated"
	"github.com/vikusku/book-freelancer/internal/gql/resolver"
	"log"
)

type Person struct {
	Name string
	Age int
}

func QueryHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		log.Println("query handler")
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func Playground(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", path)

	return func(c *gin.Context) {
		log.Println("playgrund handler")
		h.ServeHTTP(c.Writer, c.Request)
	}
}