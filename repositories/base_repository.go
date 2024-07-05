package repositories

import (
	"context"
	"fmt"

	"github.com/nanda03dev/go2ms/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *BaseRepository[T]) GetAll(ctx context.Context, filters common.Filters, sort interface{}, limit interface{}) ([]T, error) {
	filter := bson.D{}
	if filters == nil {
		filter = bson.D{}
	}

	for _, f := range filters {
		filter = append(filter, bson.E{Key: f.Key, Value: f.Value})
	}

	// Prepare options
	findOptions := options.Find()
	sortOptions := bson.D{{Key: "_id", Value: 1}}

	if sort != nil {
		fmt.Println("sort ", sort)

		temp := sort.(common.Sort)
		if temp.Order < 1 {
			temp.Order = -1
		}
		sortOptions = bson.D{{Key: temp.Key, Value: temp.Order}}
	}

	fmt.Println("sortOptions ", sortOptions)
	findOptions.SetSort(sortOptions)

	if limit != nil {
		switch v := limit.(type) {
		case int:
			findOptions.SetLimit(int64(v))
		case int64:
			findOptions.SetLimit(v)
		case *int:
			if v != nil {
				findOptions.SetLimit(int64(*v))
			}
		case *int64:
			if v != nil {
				findOptions.SetLimit(*v)
			}
		}
	}

	cursor, err := r.collection.Find(ctx, filter, findOptions)
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
