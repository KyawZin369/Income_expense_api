package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/example/graphql-mysql-api/pkg/api/graph"
	"github.com/example/graphql-mysql-api/pkg/config"
	"github.com/example/graphql-mysql-api/pkg/database"
	"github.com/example/graphql-mysql-api/pkg/middleware"
	"github.com/example/graphql-mysql-api/pkg/utils"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	// Initialize JWT
	utils.InitJWT(cfg.JWTSecret)

	r := gin.Default()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(db)}))
	
	// Add authentication middleware
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		return middleware.AuthMiddleware(ctx, nil, next)
	})

	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	if cfg.EnablePlayground {
		r.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
		})
	}

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Printf("server listening on :%s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

