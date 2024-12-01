package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createTopList(t *testing.T) *TopList {

	arg := &topListParams{
		Url:  "http://tuijian.hao1231.com/",
		Name: "热搜1",
	}

	result, err := testStore.AddTopList(context.Background(), "topList", arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Url, result.Url)
	return result
}

func TestAddTopList(t *testing.T) {
	createTopList(t)
}

func TestGetAllTopList(t *testing.T) {
	insertTopList := createTopList(t)
	filter := bson.M{}

	results, err := testStore.GetAllTopList(context.Background(), "topList", filter, 0, 3)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	var found bool
	for _, result := range results {
		if result["url"] == insertTopList.Url && result["name"] == insertTopList.Name {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted insertHotTab not found in the results")
}

func TestAddManyTopList(t *testing.T) {
	topList, err := LoadFromFile[topListParams]("../sample-data/top-list.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*topListParams

	for i := range topList {
		dummy = append(dummy, &topList[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(topList), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertTopList, err := testStore.AddManyTopList(context.Background(), "topList", dummy)

	require.NoError(t, err)
	require.Len(t, insertTopList, len(dummy))
}
