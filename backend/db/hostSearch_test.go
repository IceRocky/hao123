package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createHostSearch(t *testing.T) *HostSearch {

	arg := &HostSearchParams{
		Link:  "https://www.baidu.com1/s?tn=50000201_hao_pg&ie=utf-8&word=%E5%85%A8%E7%BD%91%E5%AF%BB%E6%89%BE%E7%9A%84%E8%B1%ABP7A525%E6%89%BE%E5%88%B0%E4%BA%86&rsv_dl=fyb_n_hao123pc",
		Title: "全网寻找的豫P7A525找到了1",
	}

	result, err := testStore.AddHostSearch(context.Background(), "hostSearch", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Link, result.Link)
	require.Equal(t, arg.Title, result.Title)
	return result
}

func TestAddHostSearch(t *testing.T) {
	createHostSearch(t)
}

func TestGetAllHostSearch(t *testing.T) {
	insertHostSearch := createHostSearch(t)
	filter := bson.M{}
	results, err := testStore.GetAllHotSearch(context.Background(), "hostSearch", filter, 0, 3)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	var found bool
	for _, result := range results {
		if result["link"] == insertHostSearch.Link && result["title"] == insertHostSearch.Title {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted host search not found in the results")
}

func TestManyHostSearch(t *testing.T) {
	hotSearches, err := LoadFromFile[HostSearchParams]("../sample-data/hot-search.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*HostSearchParams

	for i := range hotSearches {
		dummy = append(dummy, &hotSearches[i])
	}

	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(hotSearches), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertHostSearch, err := testStore.AddManyHotSearch(context.Background(), "hostSearch", dummy)

	require.NoError(t, err) // Ensure no errors occurred
	require.Len(t, insertHostSearch, len(dummy))

	for i, inserted := range insertHostSearch {
		require.NotNil(t, inserted.ID)
		require.Equal(t, dummy[i].Link, inserted.Link)
		require.Equal(t, dummy[i].Title, inserted.Title)
		require.NotNil(t, inserted.CreatedAt)
		require.NotNil(t, inserted.UpdatedAt)
	}
}
