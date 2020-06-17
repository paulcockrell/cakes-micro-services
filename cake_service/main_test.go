package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/paulcockrell/waracle-cake-service/handler"
	"github.com/paulcockrell/waracle-cake-service/repository"
	"github.com/stretchr/testify/assert"
)

const (
	dbURI      = "mongodb://localhost:27017"
	db         = "test"
	collection = "cakes"
)

var h *handler.Handler
var repo *repository.MongoRepository

func TestMain(m *testing.M) {
	repo = &repository.MongoRepository{}
	err := repo.Setup(context.TODO(), dbURI, db, collection)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Client.Disconnect(context.Background())

	// Setup request handler
	h = &handler.Handler{Repository: repo}

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	clearCakes()
	addCakes(2)

	req, _ := http.NewRequest("GET", "/cakes", nil)
	rsp := executeRequest(req)

	assert.Equal(t, http.StatusOK, rsp.Code)
}

func TestGetNonExistentCake(t *testing.T) {
	clearCakes()

	req, _ := http.NewRequest("GET", "/cakes/nonexistentid", nil)
	rsp := executeRequest(req)

	assert.Equal(t, http.StatusNotFound, rsp.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	router := setupRouter(h)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func clearCakes() {
	repo.Collection.Drop(context.TODO())
}

func addCakes(count int) {
	for i := 0; i < count; i++ {
		cake := generateCake(i)
		h.Repository.Create(context.TODO(), &cake)
	}
}

func generateCake(id int) repository.Cake {
	return repository.Cake{
		Name:     fmt.Sprintf("Cake-%d", id),
		ImageURL: fmt.Sprintf("img%d.png", id),
		Comment:  fmt.Sprintf("Comment #%d", id),
	}
}
