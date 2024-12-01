package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createAddCategory(t *testing.T) *Category {

	arg := &CategoryParams{
		Name: "综合1",
		URL:  "http://www.hao123.com/sitemap1",
		Items: []CategoryItemParams{
			{
				Name:     "高顿教育1",
				URL:      "https://www.gaodun.com1/",
				Category: func() *string { s := "综合1"; return &s }(),
			},
		},
	}
	result, err := testStore.AddCategory(context.Background(), "categories", arg)
	require.NotEmpty(t, result.Items) // Ensure items are not empty
	require.Len(t, result.Items, len(arg.Items))
	require.NoError(t, err)

	for i, item := range arg.Items {
		require.Equal(t, item.Name, result.Items[i].Name)
		require.Equal(t, item.URL, result.Items[i].URL)
		require.Equal(t, *item.Category, *result.Items[i].Category)
	}

	return result
}

func TestAddCategory(t *testing.T) {
	createAddCategory(t)
}

func TestGetAllCategory(t *testing.T) {
	insertedCategory := createAddCategory(t)

	filter := bson.M{} // Empty filter to retrieve all documents
	results, err := testStore.GetAllItemCategories(context.Background(), "categories", filter, 0, 3)

	// Step 1: General Assertions
	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Step 2: Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["name"] == insertedCategory.Name && result["url"] == insertedCategory.URL {
			// Step 3: Validate the matched category details
			require.Equal(t, insertedCategory.Name, result["name"])
			require.Equal(t, insertedCategory.URL, result["url"])

			// Step 5: Validate items if present
			if items, ok := result["items"].([]interface{}); ok {
				require.NotEmpty(t, items)

				// Validate the first item in the array
				item := items[0].(map[string]interface{})
				require.Equal(t, insertedCategory.Items[0].Name, item["name"])
				require.Equal(t, insertedCategory.Items[0].URL, item["url"])
				require.Equal(t, insertedCategory.Items[0].Category, item["category"])
			}
			found = true
			break
		}
	}

	// Step 6: Ensure the inserted category is found
	require.True(t, found, "Inserted category not found in the results")
}

func TestAddManyCategory(t *testing.T) {
	category, err := LoadFromFile[CategoryParams]("../sample-data/category.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*CategoryParams

	for i := range category {
		dummy = append(dummy, &category[i])
	}

	insertCategory, err := testStore.AddManyCategories(context.Background(), "categories", dummy)
	require.NoError(t, err) // Ensure no errors occurred
	require.NotEmpty(t, insertCategory)

	for i, inserted := range insertCategory {
		require.NotNil(t, inserted.ID)
		require.Equal(t, dummy[i].Name, inserted.Name)
		require.Equal(t, dummy[i].URL, inserted.URL)
		require.NotNil(t, inserted.CreatedAt)
		require.NotNil(t, inserted.UpdatedAt)

		// Validate items within each category, if present
		if len(dummy[i].Items) > 0 {
			require.Len(t, inserted.Items, len(dummy[i].Items))
			for j, item := range inserted.Items {
				require.Equal(t, dummy[i].Items[j].Name, item.Name)
				require.Equal(t, dummy[i].Items[j].URL, item.URL)

			}
		}
	}
}
