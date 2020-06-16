package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulcockrell/waracle-cake-service/handler"
	"github.com/paulcockrell/waracle-cake-service/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultDbURI   = "mongodb://localhost:27017"
	defaultAppPort = ":8080"
)

func main() {
	// Setup application variables
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		dbURI = defaultDbURI
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = defaultAppPort
	}

	// Setup database
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Setup request handler
	cakeCollection := client.Database("waracle").Collection("cakes")
	repo := &repository.MongoRepository{Collection: cakeCollection}
	h := &handler.Handler{Repository: repo}

	// Define API endpoints
	g := gin.Default()
	g.GET("/cakes", h.GetAll)
	g.GET("/cakes/:id", h.Get)
	g.POST("/cakes", h.Create)
	g.PUT("/cakes/:id", h.Update)
	g.DELETE("/cakes/:id", h.Delete)

	// Go!
	log.Println(fmt.Sprintf("Super dooper cake API server listening at %v", appPort))
	g.Run(appPort)
}
