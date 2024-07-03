package repositories

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type BaseRepository struct {
    collection *mongo.Collection
}

func NewBaseRepository(db *mongo.Database, collectionName string) *BaseRepository {
    return &BaseRepository{
        collection: db.Collection(collectionName),
    }
}

func (r *BaseRepository) Create(ctx context.Context, document interface{}) error {
    _, err := r.collection.InsertOne(ctx, document)
    return err
}

func (r *BaseRepository) GetAll(ctx context.Context, filter interface{}) ([]bson.M, error) {
    cursor, err := r.collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []bson.M
    for cursor.Next(ctx) {
        var result bson.M
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }

    return results, nil
}

func (r *BaseRepository) GetByID(ctx context.Context, id interface{}) (bson.M, error) {
    var result bson.M
    err := r.collection.FindOne(ctx, id).Decode(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func (r *BaseRepository) Update(ctx context.Context, filter interface{}, update interface{}) error {
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *BaseRepository) Delete(ctx context.Context, filter interface{}) error {
    _, err := r.collection.DeleteOne(ctx, filter)
    return err
}
