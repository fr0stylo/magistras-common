package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{
		collection: col,
	}
}

func (repo *Repository) GetAllPaged(ctx context.Context, query bson.M, skip int64, take int64, result interface{}) (hasNext bool, err error) {
	queryResponse, err := repo.collection.Find(ctx, query, options.Find().SetSkip(skip).SetLimit(take))
	if err != nil {
		return false, err
	}
	defer queryResponse.Close(ctx)
	hasNext = queryResponse.TryNext(ctx)

	if err := queryResponse.All(ctx, result); err != nil {
		return false, err
	}

	return
}

func (repo *Repository) GetAll(ctx context.Context, query bson.M, result interface{}) (hasNext bool, err error) {
	queryResponse, err := repo.collection.Find(ctx, query)
	if err != nil {
		return false, err
	}
	defer queryResponse.Close(ctx)
	hasNext = queryResponse.TryNext(ctx)

	if err := queryResponse.All(ctx, result); err != nil {
		return false, err
	}

	return
}

func (repo *Repository) GetById(ctx context.Context, oid string, result interface{}) (err error) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return err
	}

	return repo.Get(ctx, bson.M{"_id": id}, result)
}

func (repo *Repository) Get(ctx context.Context, query bson.M, result interface{}) (err error) {
	return repo.collection.FindOne(ctx, query).Decode(result)
}

func (repo *Repository) Insert(ctx context.Context, item interface{}) (id string, err error) {
	res, err := repo.collection.InsertOne(ctx, item)

	return res.InsertedID.(primitive.ObjectID).Hex(), err

}

func (repo *Repository) DeleteById(ctx context.Context, oid string) (err error) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return err
	}

	return repo.Delete(ctx, bson.M{"_id": id})
}

func (repo *Repository) Delete(ctx context.Context, filter interface{}) (err error) {
	number, err := repo.collection.DeleteMany(ctx, filter)

	log.Print("Deleted: ", number)
	return err
}

func (repo *Repository) Update(ctx context.Context, filter interface{}, item interface{}) error {
	_, err := repo.collection.UpdateMany(ctx, filter, item)

	return err
}

type Database interface {
	GetAllPaged(ctx context.Context, query bson.M, skip int64, take int64, result interface{}) (hasNext bool, err error)
	GetAll(ctx context.Context, query bson.M, result interface{}) (hasNext bool, err error)
	GetById(ctx context.Context, oid string, result interface{}) (err error)
	Insert(ctx context.Context, item interface{}) (id string, err error)
	Get(ctx context.Context, query bson.M, result interface{}) (err error)
	Delete(ctx context.Context, filter interface{}) (err error)
	DeleteById(ctx context.Context, oid string) (err error)
	Update(ctx context.Context, filter interface{}, item interface{}) error
}
