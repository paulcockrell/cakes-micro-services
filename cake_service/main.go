package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
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
	g.POST("/cakes", h.Create)

	// Go!
	log.Println(fmt.Sprintf("Super dooper cake API server listening at %v", appPort))
	g.Run(appPort)
}
