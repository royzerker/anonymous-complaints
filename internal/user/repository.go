package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}

type userRepositoryImpl struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		collection: db.Collection("users"),
	}
}

func (r *userRepositoryImpl) Create(user *User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *userRepositoryImpl) FindByEmail(email string) (*User, error) {
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
