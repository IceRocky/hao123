package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createTopListItems(t *testing.T) *TopListItem {
	arg := &topListItemParams{
		Url:         "http://www1.baidu.com/s?word=%E6%A8%8A%E6%8C%AF%E4%B8%9C%E8%B7%8C%E5%87%BA%E4%B8%96%E7%95%8C%E5%89%8D5%20%E7%BA%AA%E5%BD%95%E9%81%AD%E7%BB%88%E7%BB%93&tn=50000202_hao_pg&ie=utf-8&rsv_dl=fyb_n_hao123pc",
		Name:        "樊振东跌出世界前5 纪录遭终结",
		TopListName: "运动1",
	}

	result, err := testStore.AddTopListItem(context.Background(), "topListItems", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.TopListName, result.TopListName)
	require.Equal(t, arg.Url, result.Url)

	return result
}

func TestAddTopListItem(t *testing.T) {
	createTopListItems(t)
}

func TestGetAllTopListItems(t *testing.T) {
	insertTopListItems := createTopListItems(t)
	filter := bson.M{}

	results, err := testStore.GetAllTopListItem(context.Background(), "topListItems", filter, 0, 3)

	require.NoError(t, err)
	require.NotEmpty(t, results)

	var found bool
	for _, result := range results {
		if result["url"] == insertTopListItems.Url && result["topListName"] == insertTopListItems.TopListName {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted insertHotTab not found in the results")
}

func TestGetTopListItem(t *testing.T) {
	topListItem, err := LoadFromFile[topListItemParams]("../sample-data/top-list-items.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load topListItem: %v", err))
	}

	var dummy []*topListItemParams
	for i := range topListItem {
		dummy = append(dummy, &topListItem[i])
	}

	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(topListItem), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertHotTab, err := testStore.AddManyTopListItem(context.Background(), "topListItems", dummy)

	require.NoError(t, err)
	require.Len(t, insertHotTab, len(dummy))
}
