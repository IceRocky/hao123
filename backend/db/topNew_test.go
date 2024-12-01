package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createTopNews(t *testing.T) *FeedTopNew {
	arg := &feedTopNewsParams{
		Url:       "http://baijiahao.baidu.com/s?id=1816693200645667042",
		Name:      "时政纪录片丨破风逐浪启新航——习近平主席拉美之行纪实",
		FeedTitle: "推荐",
	}
	result, err := testStore.AddTopNews(context.Background(), "topNews", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Url, result.Url)
	require.Equal(t, arg.FeedTitle, result.FeedTitle)

	return result
}

func TestAddTopNews(t *testing.T) {
	createTopNews(t)
}

func TestGetAllTopNews(t *testing.T) {
	insertedTopNews := createTopNews(t) // This adds a document using AddItemCategories

	// Step 2: Retrieve all categories
	filter := bson.M{} // Empty filter to retrieve all documents
	results, err := testStore.GetAllTopNews(context.Background(), "topNews", filter, 0, 3)

	// Step 3: Assertions
	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["url"] == insertedTopNews.Url && result["name"] == insertedTopNews.Name && result["feedTitle"] == insertedTopNews.FeedTitle {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted category not found in the results")
}

func TestAddManyTopNews(t *testing.T) {
	topNews, err := LoadFromFile[feedTopNewsParams]("../sample-data/feed-top-news.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*feedTopNewsParams
	for i := range topNews {
		dummy = append(dummy, &topNews[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(topNews), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertSiteItem, err := testStore.AddManyTopNews(context.Background(), "topNews", dummy)

	require.NoError(t, err)
	require.Len(t, insertSiteItem, len(dummy))
}
