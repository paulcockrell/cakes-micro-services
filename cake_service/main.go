package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/paulcockrell/waracle-cake-service/handler"
	"github.com/paulcockrell/waracle-cake-service/repository"
)

const (
	defaultDbURI   = "mongodb://localhost:27017"
	defaultAppPort = ":8080"
)

func setupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/cakes", h.GetAll)
	r.GET("/cakes/:id", h.Get)
	r.POST("/cakes", h.Create)
	r.PUT("/cakes/:id", h.Update)
	r.DELETE("/cakes/:id", h.Delete)

	return r
}

func main() {
	ctx := context.Background()

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		dbURI = defaultDbURI
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = defaultAppPort
	}

	// Setup database
	repo := &repository.MongoRepository{}
	err := repo.Setup(ctx, dbURI, "waracle", "cakes")
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Client.Disconnect(context.Background())

	// Setup request handler
	h := &handler.Handler{Repository: repo}

	// Setup router
	r := setupRouter(h)

	// Go!
	log.Println(fmt.Sprintf("Cake API server listening at %v", appPort))
	r.Run(appPort)
}
