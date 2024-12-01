package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createHotList(t *testing.T) *HotList {
	arg := &hotListParams{
		HotlistTab: "电影",
		URL:        "https://www.baidu.com/s?wd=%E8%80%81%E5%B8%88%EF%BC%8C%E5%88%AB%E5%93%AD+%E7%94%B5%E5%BD%B1&sa=fyb_hp_movie_hao123&rsv_dl=fyb_hp_movie_hao123&from=hao123&tn=50000181_hao_pg",
		Name:       "老师，别哭",
		ImageLink:  "https://fyb-2.cdn.bcebos.com/hotboard_image/65075a4fe4762455a85737650fa035b8",
		InfoTexts: []string{
			"热搜指数 : 123439",
			"地区 : 中国大陆",
			"类型 : 剧情"}}

	result, err := testStore.AddHotList(context.Background(), "hotLists", arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.HotlistTab, result.HotlistTab)
	require.Equal(t, arg.URL, result.URL)
	require.Equal(t, arg.ImageLink, result.ImageLink)

	return result
}

func TestAddHotList(t *testing.T) {
	createHotList(t)
}

func TestGetAllHotLists(t *testing.T) {
	insertedHotLists := createHotList(t)
	filter := bson.M{}

	results, err := testStore.GetAllHotList(context.Background(), "hotLists", filter, 0, 3)

	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["url"] == insertedHotLists.URL && result["name"] == insertedHotLists.Name && result["imageLink"] == insertedHotLists.ImageLink {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted category not found in the results")
}

func TestAddManyHotLists(t *testing.T) {
	hotLists, err := LoadFromFile[hotListParams]("../sample-data/hotlist-items.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*hotListParams
	for i := range hotLists {
		dummy = append(dummy, &hotLists[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(hotLists), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertSiteItem, err := testStore.AddManyHotList(context.Background(), "hotLists", dummy)

	require.NoError(t, err)
	require.Len(t, insertSiteItem, len(dummy))
}
