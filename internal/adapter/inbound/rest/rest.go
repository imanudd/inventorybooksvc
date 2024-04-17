package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/imanudd/inventorybooksvc/config"
	_ "github.com/imanudd/inventorybooksvc/docs"
	"github.com/imanudd/inventorybooksvc/internal/adapter/inbound/rest/handler"
	"github.com/imanudd/inventorybooksvc/internal/adapter/inbound/rest/handler/middleware"
	inregistry "github.com/imanudd/inventorybooksvc/internal/core/port/inbound/registry"
	outregistry "github.com/imanudd/inventorybooksvc/internal/core/port/outbound/registry"
)

type Server struct {
	app    *gin.Engine
	config *config.MainConfig
}

// NewRest
// @title Inventory Service API
// @version 1.0
// @description Inventory Service API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name authorization

func New(config *config.MainConfig) *Server {
	if config.Environment != "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	})

	return &Server{
		app:    app,
		config: config,
	}
}

func (s *Server) RegisterHandler(cfg *config.MainConfig, repo outregistry.RepositoryRegistry, service inregistry.ServiceRegistry) error {
	s.app.Use(gin.Recovery())

	auth := middleware.NewAuthMiddleware(cfg, repo.GetUserRepository())

	handler := handler.NewHandler(service)
	inventorySvc := s.app.Group("/inventorysvc")

	inventorySvc.GET("/", func(c *gin.Context) {
		c.JSON(200, "welcome to inventory service")
	})

	inventorySvc.POST("/auth/register", handler.Register)
	inventorySvc.POST("/auth/login", handler.Login)

	inventorySvc.POST("/managements/book", auth.JWTAuth(handler.AddBook))
	inventorySvc.PUT("/managements/book/:id", auth.JWTAuth(handler.UpdateBook))
	inventorySvc.DELETE("/managements/book/:id", auth.JWTAuth(handler.DeleteBook))
	inventorySvc.GET("/managements/book/:id", auth.JWTAuth(handler.GetDetailBook))

	inventorySvc.POST("/managements/author", auth.JWTAuth(handler.CreateAuthor))
	inventorySvc.POST("managements/author/:id", auth.JWTAuth(handler.AddAuthorBook))
	inventorySvc.GET("managements/author/:id/list", auth.JWTAuth(handler.GetListBookByAuthor))
	inventorySvc.DELETE("managements/author/:id/books/:bookid", auth.JWTAuth(handler.DeleteBookByAuthor))

	return nil
}

func (s *Server) Serve() (err error) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.ServicePort),
		Handler: s.app,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error: %s\n", err)
		}
	}()

	log.Println("-------------------------------------------")
	log.Println("server started")
	log.Printf("running on port %d\n", s.config.ServicePort)
	log.Println("-------------------------------------------")

	return gracefulShutdown(server)
}

func gracefulShutdown(srv *http.Server) error {
	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Println("Shutting down server...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error while shutting down Server. Initiating force shutdown...")
		return err
	}

	log.Println("Server exiting...")

	return nil
}
