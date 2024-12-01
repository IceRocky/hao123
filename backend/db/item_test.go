package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createAddItem(t *testing.T) *Item {

	arg := &ItemParams{
		Name:     "巨龙猎手1",
		URL:      "https://wan.baidu.com1/cover?gameId=58573089&idfrom=4084",
		Category: func() *string { s := "页游"; return &s }(),
	}

	result, err := testStore.AddItem(context.Background(), "item", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.URL, result.URL)
	require.Equal(t, arg.Category, result.Category)
	return result

}

func TestAddItem(t *testing.T) {
	createAddItem(t)
}

func TestGetAllItem(t *testing.T) {
	// Step 1: Add a test category
	insertedCategory := createAddItem(t) // This adds a document using AddItemCategories

	// Step 2: Retrieve all categories
	filter := bson.M{} // Empty filter to retrieve all documents
	results, err := testStore.GetAllItem(context.Background(), "item", filter, 0, 3)

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

	require.True(t, found, "Inserted ITEM not found in the results")
}

func TestAddManyItem(t *testing.T) {
	items, err := LoadFromFile[ItemParams]("../sample-data/items.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	// Print loaded categories

	var dummy []*ItemParams
	for i := range items {
		dummy = append(dummy, &items[i])

	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(items), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	inserteIem, err := testStore.AddManyItem(context.Background(), "item", dummy)
	require.NoError(t, err)
	require.Len(t, inserteIem, len(dummy))
}
