package main

import (
	"bytes"
	"context"
	"encoding/json"
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

func TestGetCake(t *testing.T) {
	clearCakes()
	cakes := addCakes(1)

	url := fmt.Sprintf("/cakes/%s", cakes[0].ID.Hex())
	req, _ := http.NewRequest("GET", url, nil)
	rsp := executeRequest(req)

	var cake repository.Cake
	json.Unmarshal(rsp.Body.Bytes(), &cake)

	assert.Equal(t, http.StatusOK, rsp.Code)
	assert.Equal(t, cakes[0], &cake)

}

func TestCreateCake(t *testing.T) {
	clearCakes()

	var jsonStr = []byte(`{"name":"Cake","comment":"Comment","image_url":"ImgURL","yum_factor":10}`)
	req, _ := http.NewRequest("POST", "/cakes", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rsp := executeRequest(req)

	assert.Equal(t, http.StatusOK, rsp.Code)

	var cake repository.Cake
	json.Unmarshal(rsp.Body.Bytes(), &cake)

	assert.NotNil(t, cake.ID)
	assert.Equal(t, "Cake", cake.Name)
	assert.Equal(t, "Comment", cake.Comment)
	assert.Equal(t, "ImgURL", cake.ImageURL)
	assert.Equal(t, int8(10), cake.YumFactor)
}

func TestUpdateCake(t *testing.T) {
	clearCakes()
	cakes := addCakes(1)

	// Update cake
	var jsonStr = []byte(`{"name":"Updated name","comment":"Updated comment","image_url":"Updated ImgURL","yum_factor":1}`)
	url := fmt.Sprintf("/cakes/%s", cakes[0].ID.Hex())
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rsp := executeRequest(req)

	assert.Equal(t, http.StatusOK, rsp.Code)

	// Get updated cake
	req, _ = http.NewRequest("GET", url, nil)
	rsp = executeRequest(req)

	var cake repository.Cake
	json.Unmarshal(rsp.Body.Bytes(), &cake)

	assert.Equal(t, cakes[0].ID, cake.ID)
	assert.Equal(t, "Updated name", cake.Name)
	assert.Equal(t, "Updated comment", cake.Comment)
	assert.Equal(t, "Updated ImgURL", cake.ImageURL)
	assert.Equal(t, int8(1), cake.YumFactor)
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

func addCakes(count int) []*repository.Cake {
	var cakes []*repository.Cake

	for i := 0; i < count; i++ {
		cake := generateCake(i)
		h.Repository.Create(context.TODO(), &cake)
		cakes = append(cakes, &cake)
	}

	return cakes
}

func generateCake(id int) repository.Cake {
	return repository.Cake{
		Name:     fmt.Sprintf("Cake-%d", id),
		ImageURL: fmt.Sprintf("img%d.png", id),
		Comment:  fmt.Sprintf("Comment #%d", id),
	}
}
