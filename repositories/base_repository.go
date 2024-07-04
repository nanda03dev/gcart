package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository[T any] struct {
	collection *mongo.Collection
}

func NewBaseRepository[T any](db *mongo.Database, collectionName string) *BaseRepository[T] {
	return &BaseRepository[T]{
		collection: db.Collection(collectionName),
	}
}

func (r *BaseRepository[T]) Create(ctx context.Context, document T) error {
	_, err := r.collection.InsertOne(ctx, document)
	return err
}

func (r *BaseRepository[T]) GetAll(ctx context.Context, filter interface{}) ([]T, error) {
	if filter == nil {
		filter = bson.D{}
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []T
	for cursor.Next(ctx) {
		var result T
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *BaseRepository[T]) GetByID(ctx context.Context, id primitive.ObjectID) (T, error) {
	var result T
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	return result, err
}

func (r *BaseRepository[T]) Update(ctx context.Context, id primitive.ObjectID, update T) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
