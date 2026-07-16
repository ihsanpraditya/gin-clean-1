package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/casbin/casbin/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ihsanpraditya/gin-clean-1/graph"
	"github.com/ihsanpraditya/gin-clean-1/internal/database"
	"github.com/ihsanpraditya/gin-clean-1/internal/middleware"
	"github.com/ihsanpraditya/gin-clean-1/internal/repository"
	"github.com/ihsanpraditya/gin-clean-1/internal/service"
)

func SetupRouter(r *gin.Engine) {
	userRepo := repository.NewUserRepository(database.DB)
	roleRepo := repository.NewRoleRepository(database.DB)
	
	userSvc := service.NewUserService(userRepo)
	roleSvc := service.NewRoleService(roleRepo)
	authSvc := service.NewAuthService(userRepo)
	
	val := validator.New()

	// Initialize Casbin Enforcer using the config files
	enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	config := graph.Config{
		Resolvers: &graph.Resolver{
			UserSvc: userSvc,
			RoleSvc: roleSvc,
			AuthSvc: authSvc,
			Validator: val,
		},
	}

	setupAuthentication(&config)
	setupAuthorization(&config, userSvc, enforcer)

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

func setupAuthentication(config *graph.Config) {
	// Protect using GraphQL Directives + Casbin
	config.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		_, ok := ctx.Value(middleware.UserContextKey).(uint)
		if !ok {
			return nil, errors.New("unauthorized: please login")
		}

		return next(ctx)
	}
}

func setupAuthorization(config *graph.Config, userSvc *service.UserService, enforcer *casbin.Enforcer) {
	// Setup new directive for authorization
	config.Directives.Can = func(ctx context.Context, obj interface{}, next graphql.Resolver, resource string, action string) (interface{}, error) {
		userID, ok := ctx.Value(middleware.UserContextKey).(uint)
		if !ok {
			return nil, errors.New("unauthorized: please login")
		}

		user, err := userSvc.GetUserByID(ctx, userID)
		if err != nil {
			return nil, errors.New("unauthorized: user record not found")
		}

		allowed := false
		for _, role := range user.Roles {
			ok, err := enforcer.Enforce(role.Name, resource, action)
			if err == nil && ok {
				allowed = true
				break
			}
		}

		if !allowed {
			return nil, fmt.Errorf("forbidden: you do not have permission to perform '%s' on '%s'", action, resource)
		}

		return next(ctx)
	}
	
}