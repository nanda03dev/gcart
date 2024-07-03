package repositories

import (
	"context"

	"github.com/nanda03dev/oms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetByID(id primitive.ObjectID) (models.User, error)
	Update(user models.User) error
	Delete(id primitive.ObjectID) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) Create(user models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *userRepository) GetAll() ([]models.User, error) {
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, cursor.Err()
}

func (r *userRepository) GetByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *userRepository) Update(user models.User) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *userRepository) Delete(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
