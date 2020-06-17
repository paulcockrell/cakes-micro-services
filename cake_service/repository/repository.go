package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cake - cake definition
type Cake struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Comment   string             `json:"comment" bson:"comment"`
	ImageURL  string             `json:"imageUrl" bson:"imageUrl"`
	YumFactor int8               `json:"yumFactor" bson:"yumFactor"`
}

type Repository interface {
	Create(ctx context.Context, cake *Cake) error
	GetAll(ctx context.Context) ([]*Cake, error)
	Get(ctx context.Context, id string) (*Cake, error)
	Update(ctx context.Context, id string, cake *Cake) (*Cake, error)
	Delete(ctx context.Context, id string) error
}

type MongoRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// Setup - Establish connection with MonogoDB
func (r *MongoRepository) Setup(ctx context.Context, dbURI, db, collection string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	r.Client = client
	r.Collection = client.Database(db).Collection(collection)

	return nil
}

// Create - Create a new cake record
func (r *MongoRepository) Create(ctx context.Context, cake *Cake) error {
	res, err := r.Collection.InsertOne(ctx, cake)
	if err != nil {
		return err
	}

	cake.ID = res.InsertedID.(primitive.ObjectID)
	return nil
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

// Get - Get on cake by ID
func (r *MongoRepository) Get(ctx context.Context, id string) (*Cake, error) {
	var cake Cake
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": docID}
	err = r.Collection.FindOne(ctx, filter).Decode(&cake)
	if err != nil {
		return nil, err
	}

	return &cake, nil
}

// Update - Update cake
func (r *MongoRepository) Update(ctx context.Context, id string, cake *Cake) (*Cake, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := r.Collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{
				{"name", cake.Name},
				{"yumFactor", cake.YumFactor},
				{"comment", cake.Comment},
				{"imageUrl", cake.ImageURL},
			}},
		},
	)
	if err != nil {
		return nil, err
	}
	if res.ModifiedCount < 1 {
		return nil, errors.New("No record found to update, or nothing to update")
	}

	cake.ID = docID
	return cake, nil
}

// Delete - Delete cake by ID
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		return err
	}

	return nil
}
