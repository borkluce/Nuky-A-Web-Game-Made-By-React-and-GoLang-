package repo

import (
	"context"
	"services/internal/province/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// UpdateDestroymentRoundOfTheWorstProvince finds the province with the highest (attackCount - supportCount)
// and sets its destroymentRound to the given round count
func (pr *ProvinceRepo) UpdateDestroymentRoundOfTheWorstProvince(ctx context.Context, roundCount int) error {
	// Create aggregation pipeline to find the province with highest score difference
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
		{
			"$limit": 1, // Get only the worst province
		},
	}

	// Execute the aggregation to find the worst province
	cursor, err := pr.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Get the worst province
	var worstProvince model.Province
	if cursor.Next(ctx) {
		if err := cursor.Decode(&worstProvince); err != nil {
			return err
		}
	} else {
		return nil
	}

	// Update the destroyment round of the worst province
	filter := bson.M{"_id": worstProvince.ID}
	update := bson.M{
		"$set": bson.M{"destroymentRound": roundCount},
	}

	_, err = pr.collection.UpdateOne(ctx, filter, update)
	return err
}

// ResetAllProvinceCounts resets attackCount and supportCount to 0 for all provinces
func (pr *ProvinceRepo) ResetAllProvinceCounts(ctx context.Context) error {
	filter := bson.M{} // Empty filter to match all documents
	update := bson.M{
		"$set": bson.M{
			"attackCount":  0,
			"supportCount": 0,
		},
	}

	// Use UpdateMany to update all provinces
	_, err := pr.collection.UpdateMany(ctx, filter, update)
	return err
}
