package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"services/internal/province/model"
)

type ProvinceRepo struct {
	collection *mongo.Collection
}

// NewProvinceRepo creates a new province repository
func NewProvinceRepo(collection *mongo.Collection) *ProvinceRepo {
	return &ProvinceRepo{
		collection: collection,
	}
}

// GetAll retrieves all provinces from the database
func (pr *ProvinceRepo) GetAll(ctx context.Context) ([]model.Province, error) {
	var provinces []model.Province
	cursor, err := pr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var p model.Province

		// Raw data
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		provinces = append(provinces, p)
	}
	return provinces, nil
}

// UpdateProvinceByID updates the attack or support count for a province
func (pr *ProvinceRepo) UpdateProvinceByID(ctx context.Context, id string, isAttackNorSupport bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var update bson.M
	if isAttackNorSupport {
		update = bson.M{
			"$inc": bson.M{"attackCount": 1},
		}
	} else {
		update = bson.M{
			"$inc": bson.M{"supportCount": 1},
		}
	}

	filter := bson.M{"_id": objectID}

	_, err = pr.collection.UpdateOne(ctx, filter, update)
	// return error if ID format is invalid
	return err
}

// GetProvincesByScoreDifference retrieves provinces sorted by the difference between attackCount and supportCount
func (pr *ProvinceRepo) GetProvincesByScoreDifference(ctx context.Context) ([]model.Province, error) {
	// Create a pipeline to calculate difference and sort
	pipeline := []bson.M{
		{
			"$addFields": bson.M{
				"scoreDifference": bson.M{
					"$subtract": []string{"$attackCount", "$supportCount"},
				},
			},
		},
		{
			"$sort": bson.M{"scoreDifference": -1}, // Sort by difference in descending order
		},
	}

	// Execute the aggregation
	cursor, err := pr.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate through the results
	var provinces []model.Province
	for cursor.Next(ctx) {
		var p model.Province
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		provinces = append(provinces, p)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return provinces, nil
}
