package router

import (
	"context"
	"net/http"
	"os"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/ihsanpraditya/gin-clean-1/graph"
	"github.com/ihsanpraditya/gin-clean-1/internal/database"
	"github.com/ihsanpraditya/gin-clean-1/internal/repository"
	"github.com/ihsanpraditya/gin-clean-1/internal/service"
	"github.com/ihsanpraditya/gin-clean-1/internal/middleware"
)

func SetupRouter(r *gin.Engine) {
	userRepo := repository.NewUserRepository(database.DB)
	userSvc := service.NewUserService(userRepo)

	config := graph.Config{
		Resolvers: &graph.Resolver{
			UserSvc: userSvc,
		},
	}

	config.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		// Ambil value dari context yang di-inject oleh Gin AuthMiddleware sebelumnya
		_, ok := ctx.Value(middleware.UserContextKey).(uint)
		if !ok {
			return nil, errors.New("access denied: unauthorized")
		}

		// Jika lolos cek, lanjutkan eksekusi ke resolver asli
		return next(ctx)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	secretKey := []byte(os.Getenv("JWT_KEY"))
	r.Use(middleware.AuthMiddleware(secretKey))

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
