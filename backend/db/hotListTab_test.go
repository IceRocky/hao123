package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createHotTabs(t *testing.T) *HotListTab {
	arg := &hotTabParams{
		Url:  "https://top.baidu.com/board?tab=movie&sa=fyb_movie_hao123",
		Name: "电影",
	}

	result, err := testStore.AddHotTab(context.Background(), "hotTab", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Url, result.Url)

	return result
}

func TestAddHotTab(t *testing.T) {
	createHotTabs(t)
}

func TestGetAllHotTab(t *testing.T) {
	insertHotTab := createHotTabs(t)
	filter := bson.M{}

	results, err := testStore.GetAllHotTab(context.Background(), "hotTab", filter, 0, 3)

	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["url"] == insertHotTab.Url && result["name"] == insertHotTab.Name {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted insertHotTab not found in the results")
}

func TestAddManyHotTab(t *testing.T) {
	hotTab, err := LoadFromFile[hotTabParams]("../sample-data/hotlist-tab.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load hotListTab: %v", err))
	}

	var dummy []*hotTabParams
	for i := range hotTab {
		dummy = append(dummy, &hotTab[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(hotTab), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertHotTab, err := testStore.AddManyHotTab(context.Background(), "hotTab", dummy)

	require.NoError(t, err)
	require.Len(t, insertHotTab, len(dummy))
}
