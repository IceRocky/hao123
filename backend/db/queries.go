package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Queries interface {
	AddItemCategories(ctx context.Context, collectionName string, itemCategories *AddCategoryParams) (*ItemCategories, error)
	GetAllItemCategories(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddManyItemCategories(ctx context.Context, collectionName string, itemCategories []*AddCategoryParams) ([]*ItemCategories, error)
	GetAllCategories(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddCategory(ctx context.Context, collectionName string, category *CategoryParams) (*Category, error)
	AddManyCategories(ctx context.Context, collectionName string, categories []*CategoryParams) ([]*Category, error)
	GetAllHotSearch(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddHostSearch(ctx context.Context, collectionName string, hostSearch *HostSearchParams) (*HostSearch, error)
	AddManyHotSearch(ctx context.Context, collectionName string, hostSearch []*HostSearchParams) ([]*HostSearch, error)
	GetAllItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddItem(ctx context.Context, collectionName string, item *ItemParams) (*Item, error)
	AddManyItem(ctx context.Context, collectionName string, items []*ItemParams) ([]*Item, error)
	GetAllSiteItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddSiteItem(ctx context.Context, collectionName string, siteItem *siteItemsParams) (*SiteItems, error)
	AddManySiteItem(ctx context.Context, collectionName string, siteItems []*siteItemsParams) ([]*SiteItems, error)
	GetAllCommonSiteItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddCommonSiteItem(ctx context.Context, collectionName string, commonSiteItem *commonSiteParams) (*CommonSite, error)
	AddManyCommonSiteItem(ctx context.Context, collectionName string, commonSiteItems []*commonSiteParams) ([]*CommonSite, error)
	GetAllTopNews(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddTopNews(ctx context.Context, collectionName string, feedTopNew *feedTopNewsParams) (*FeedTopNew, error)
	AddManyTopNews(ctx context.Context, collectionName string, feedTopNew []*feedTopNewsParams) ([]*FeedTopNew, error)
	GetAllGovSites(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddGovSite(ctx context.Context, collectionName string, govSites *govSitesParams) (*GovSites, error)
	AddManyGovSites(ctx context.Context, collectionName string, govSites []*govSitesParams) ([]*GovSites, error)
	GetAllHotList(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddHotList(ctx context.Context, collectionName string, hotList *hotListParams) (*HotList, error)
	AddManyHotList(ctx context.Context, collectionName string, hotList []*hotListParams) ([]*HotList, error)
	GetAllHotTab(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddHotTab(ctx context.Context, collectionName string, hotTab *hotTabParams) (*HotListTab, error)
	AddManyHotTab(ctx context.Context, collectionName string, hotTab []*hotTabParams) ([]*HotListTab, error)
	GetAllTopList(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddTopList(ctx context.Context, collectionName string, topList *topListParams) (*TopList, error)
	AddManyTopList(ctx context.Context, collectionName string, topList []*topListParams) ([]*TopList, error)
	GetAllTopListItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error)
	AddTopListItem(ctx context.Context, collectionName string, topListItem *topListItemParams) (*TopListItem, error)
	AddManyTopListItem(ctx context.Context, collectionName string, topListItems []*topListItemParams) ([]*TopListItem, error)
}

// 金靖预言了麦学
type MongoQueries struct {
	DB *mongo.Database
}

func NewMongoQueries(db *mongo.Database) Queries {
	return &MongoQueries{
		DB: db,
	}
}

// ExecuteQuery performs generic operations on any collection
func (mq *MongoQueries) ExecuteQuery(ctx context.Context, collectionName string, queryFunc func(*mongo.Collection) error) error {

	collection := mq.DB.Collection(collectionName)
	return queryFunc(collection)
}

func LoadFromFile[T any](filePath string) ([]T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the JSON file into a slice of the generic type T
	var data []T
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return data, nil
}
