package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createSiteItem(t *testing.T) *SiteItems {

	arg := &siteItemsParams{
		Name: "百度",
		Icon: "https://dgss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2021-3-5/baidulogo.png",
		Link: "http://www.baidu.com/?tn=sitehao123_15",
	}
	result, err := testStore.AddSiteItem(context.Background(), "siteItem", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Icon, result.Icon)
	require.Equal(t, arg.Link, result.Link)

	return result
}

func TestAddSiteItem(t *testing.T) {
	createSiteItem(t)
}

func TestGetAllSiteItem(t *testing.T) {
	// Step 1: Add a test category
	insertedCategory := createSiteItem(t) // This adds a document using AddItemCategories

	// Step 2: Retrieve all categories
	filter := bson.M{} // Empty filter to retrieve all documents
	results, err := testStore.GetAllSiteItem(context.Background(), "siteItem", filter, 0, 3)

	// Step 3: Assertions
	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["name"] == insertedCategory.Name && result["icon"] == insertedCategory.Icon && result["link"] == insertedCategory.Link {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted category not found in the results")
}

func TestAddManySiteItem(t *testing.T) {
	siteItems, err := LoadFromFile[siteItemsParams]("../sample-data/site-items.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*siteItemsParams
	for i := range siteItems {
		dummy = append(dummy, &siteItems[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(siteItems), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertSiteItem, err := testStore.AddManySiteItem(context.Background(), "siteItem", dummy)

	require.NoError(t, err)
	require.Len(t, insertSiteItem, len(dummy))
}
