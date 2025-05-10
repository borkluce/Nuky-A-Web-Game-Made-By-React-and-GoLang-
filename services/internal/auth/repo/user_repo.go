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

func NewUserRepo(collection *mongo.Collection) *UserRepo {
	return &UserRepo{
		collection: collection,
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
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"email":        user.Email,
			"password":     user.Password,
			"lastMoveDate": user.LastMoveDate,
		},
	}

	_, err := ur.collection.UpdateOne(ctx, filter, update)
	return err
}

func (ur *UserRepo) CreateUser(ctx context.Context, user model.User) (primitive.ObjectID, error) {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	result, err := ur.collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, nil
}

func (ur *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	filter := bson.M{"email": email}

	if err := ur.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	filter := bson.M{"username": username}

	if err := ur.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
