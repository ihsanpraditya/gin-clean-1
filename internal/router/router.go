package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/ihsanpraditya/gin-clean-1/graph"
	"github.com/ihsanpraditya/gin-clean-1/internal/database"
	"github.com/ihsanpraditya/gin-clean-1/internal/repository"
	"github.com/ihsanpraditya/gin-clean-1/internal/service"
)

func SetupRouter(r *gin.Engine) {
	userRepo := repository.NewUserRepository(database.DB)
	userSvc := service.NewUserService(userRepo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UserSvc: userSvc,
		},
	}))

	// NOTE: apply middleware here?

	// Playground untuk testing via UI browser
	r.GET("/", ginHandler(playground.Handler("GraphQL Playground", "/query")))
	r.POST("/query", ginHandler(srv))
}

// ginHandler adalah fungsi jembatan agar handler standard http.Handler milik gqlgen cocok dengan Gin
func ginHandler(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
