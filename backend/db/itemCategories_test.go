package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createAddItemCategories(t *testing.T) *ItemCategories {
	arg := &AddCategoryParams{

		Name: "综合1",
		URL:  "http://www.hao123.com/sitemap1",
	}
	result, err := testStore.AddItemCategories(context.Background(), "itemCategories", arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.URL, result.URL)
	return result
}

func TestAddItemCategories(t *testing.T) {
	createAddItemCategories(t)
}

func TestGetAllItemCategories(t *testing.T) {
	// Step 1: Add a test category
	insertedCategory := createAddItemCategories(t) // This adds a document using AddItemCategories

	// Step 2: Retrieve all categories
	filter := bson.M{} // Empty filter to retrieve all documents
	results, err := testStore.GetAllItemCategories(context.Background(), "itemCategories", filter, 0, 3)

	// Step 3: Assertions
	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["name"] == insertedCategory.Name && result["url"] == insertedCategory.URL {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted category not found in the results")
}

func TestAddManyItemCategories(t *testing.T) {
	var dummy []*AddCategoryParams

	itemCategories, err := LoadFromFile[AddCategoryParams]("../sample-data/itemCategories.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	for i := range itemCategories {
		dummy = append(dummy, &itemCategories[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(itemCategories), len(dummy), "Dummy slice should contain the same number of items as the loaded data")
	insertedCategories, err := testStore.AddManyItemCategories(context.Background(), "itemCategories", dummy)
	require.NoError(t, err) // Ensure no errors occurred
	require.Len(t, insertedCategories, len(dummy))

	// Fetch the database reference via the store
	//	db := testStore.GetDatabase()
	//	require.NotNil(t, db)
	//
	//	// Clean up the collection
	//	err = db.Collection("ItemCategories").Drop(context.Background())
	//	require.NoError(t, err)

}
