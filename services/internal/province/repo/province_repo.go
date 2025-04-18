package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"services/internal/province/model"
)

type ProvinceRepo struct {
	collection *mongo.Collection
}

func NewProvinceRepo(db *mongo.Database) *ProvinceRepo {
	return &ProvinceRepo{
		collection: db.Collection("provinces"),
	}
}

func (pr *ProvinceRepo) GetAll(ctx contex.Context) ([]model.Province, error) {
	var provinces []model.Province
	cursor, err := pr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var p model.Province
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		provinces = append(provinces, p)
	}
	return provinces, nil
}

func (pr *ProvinceRepo) UpdateProvince(ctx context.Context, id int, isAttackNorSupport bool) error {

}
