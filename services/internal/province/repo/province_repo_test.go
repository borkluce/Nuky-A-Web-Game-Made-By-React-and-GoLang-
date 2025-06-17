package repo

/*
import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"services/internal/province/model"
)

func setupTestDB(t *testing.T) (*mongo.Database, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use only the first 10 characters of the ObjectID
	shortID := primitive.NewObjectID().Hex()[:10]
	dbName := "test_" + shortID

	uri := "mongodb+srv://vafaill_master:jqSqKsjGHSa8hlK1@cluster0.uuzqws6.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB Atlas: %v", err)
	}

	// Return the database and a cleanup function
	return client.Database(dbName), func() {
		if err := client.Database(dbName).Drop(context.Background()); err != nil {
			t.Logf("Failed to drop test database: %v", err)
		}
		if err := client.Disconnect(context.Background()); err != nil {
			t.Logf("Failed to disconnect MongoDB client: %v", err)
		}
	}
}

func TestNewProvinceRepo(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewProvinceRepo(db)
	assert.NotNil(t, repo)
	assert.Equal(t, "provinces", repo.collection.Name())
}

func TestGetAll(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create test data
	provinceRepo := NewProvinceRepo(db)
	testProvinces := []interface{}{
		model.Province{
			ID:           primitive.NewObjectID(),
			ProvinceName: "Zartistan",
			AttackCount:  5,
			SupportCount: 3,
		},
		model.Province{
			ID:           primitive.NewObjectID(),
			ProvinceName: "Zortistan",
			AttackCount:  2,
			SupportCount: 7,
		},
	}

	// Insert test data
	ctx := context.Background()
	_, err := provinceRepo.collection.InsertMany(ctx, testProvinces)
	assert.NoError(t, err)

	// Test GetAll method
	provinces, err := provinceRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, provinces, 2)

	// Verify the data
	assert.Contains(t, []string{provinces[0].ProvinceName, provinces[1].ProvinceName}, "Zartistan")
	assert.Contains(t, []string{provinces[0].ProvinceName, provinces[1].ProvinceName}, "Zortistan")
}

func TestUpdateProvinceByID(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	provinceRepo := NewProvinceRepo(db)
	ctx := context.Background()

	// Create a test province
	provinceID := primitive.NewObjectID()
	testProvince := model.Province{
		ID:           provinceID,
		ProvinceName: "Test Province",
		AttackCount:  0,
		SupportCount: 0,
	}

	_, err := provinceRepo.collection.InsertOne(ctx, testProvince)
	assert.NoError(t, err)

	// Test 1: Increment attack count
	err = provinceRepo.UpdateProvinceByID(ctx, provinceID.Hex(), true)
	assert.NoError(t, err)

	// Verify the update
	var updatedProvince model.Province
	err = provinceRepo.collection.FindOne(ctx, bson.M{"_id": provinceID}).Decode(&updatedProvince)
	assert.NoError(t, err)
	assert.Equal(t, 1, updatedProvince.AttackCount)
	assert.Equal(t, 0, updatedProvince.SupportCount)

	// Test 2: Increment support count
	err = provinceRepo.UpdateProvinceByID(ctx, provinceID.Hex(), false)
	assert.NoError(t, err)

	// Verify the update
	err = provinceRepo.collection.FindOne(ctx, bson.M{"_id": provinceID}).Decode(&updatedProvince)
	assert.NoError(t, err)
	assert.Equal(t, 1, updatedProvince.AttackCount)
	assert.Equal(t, 1, updatedProvince.SupportCount)

	// Test 3: Invalid ID
	err = provinceRepo.UpdateProvinceByID(ctx, "invalid-id", true)
	assert.Error(t, err)
}
*/
