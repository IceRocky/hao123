package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createCommonSite(t *testing.T) *CommonSite {
	arg := &commonSiteParams{
		Name: "网易云音乐",
		Url:  "https://music.163.com",
		Icon: "",
	}

	result, err := testStore.AddCommonSiteItem(context.Background(), "commonSite", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Url, result.Url)

	return result
}

func TestAddCommonSite(t *testing.T) {
	createCommonSite(t)
}

func TestGetAllCommonSite(t *testing.T) {
	insertedCommonSite := createCommonSite(t)

	filter := bson.M{} // Empty filter to retrieve all documents
	results, err := testStore.GetAllCommonSiteItem(context.Background(), "commonSite", filter, 0, 3)

	// Step 3: Assertions
	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		fmt.Println("1", result["url"])
		fmt.Println("2", insertedCommonSite.Url)
		if result["name"] == insertedCommonSite.Name && result["url"] == insertedCommonSite.Url {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted common site item not found in the results")
}

func TestAddMannyCommonSite(t *testing.T) {
	commonSite, err := LoadFromFile[commonSiteParams]("../sample-data/common-site.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*commonSiteParams

	for i := range commonSite {
		dummy = append(dummy, &commonSite[i])
	}

	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(commonSite), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertedCommonSite, err := testStore.AddManyCommonSiteItem(context.Background(), "commonSite", dummy)

	require.NoError(t, err)
	require.Len(t, insertedCommonSite, len(dummy))
}
