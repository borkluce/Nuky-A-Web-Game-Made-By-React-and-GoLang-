package repo

import (
	"context"
	"services/internal/auth/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		collection: db.Collection("users"),
	}
}

func (ur *UserRepo) GetUserByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	var user *model.User
	filter := bson.M{"_id": id}

	if err := ur.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepo) PutUser(ctx context.Context, user model.User) error {
	filter := bson.M{"_id", user.ID}
	update := bson.M{
		"$set": bson.M{
			"email":        user.Email,
			"password":     user.Password,
			"lastMoveDate": user.LastMoveDate,
		},
	}
}
