package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, cake *Cake) error
	GetAll(ctx context.Context) ([]*Cake, error)
}

type MongoRepository struct {
	Collection *mongo.Collection
}

// Cake - cake definition
type Cake struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	ImageURL  string `json:"image_url"`
	YumFactor int8   `json:"yum_factor"`
}

// Create - Create a new cake record
func (r *MongoRepository) Create(ctx context.Context, cake *Cake) error {
	_, err := r.Collection.InsertOne(ctx, cake)
	return err
}

// GetAll - Get all cake records
func (r *MongoRepository) GetAll(ctx context.Context) ([]*Cake, error) {
	cur, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var cakes []*Cake
	for cur.Next(ctx) {
		var cake *Cake
		if err := cur.Decode(&cake); err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}
