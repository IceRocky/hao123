package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type hotListParams struct {
	HotlistTab string   `bson:"hotlistTab" json:"hotlistTab"`
	URL        string   `bson:"url" json:"url"`
	Name       string   `bson:"name" json:"name"`
	ImageLink  string   `bson:"imageLink" json:"imageLink"`
	InfoTexts  []string `bson:"infoTexts" json:"infoTexts"`
}

func (mq *MongoQueries) GetAllHotList(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
	if filter == nil {
		filter = bson.M{}
	}
	var results []bson.M

	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		findOptions := options.Find().SetSkip(skip).SetLimit(limit)
		cursor, err := collection.Find(ctx, filter, findOptions)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var document bson.M
			if err := cursor.Decode(&document); err != nil {
				return err
			}
			results = append(results, document)
		}

		if err := cursor.Err(); err != nil {
			return err
		}

		return nil
	})

	return results, err
}

func (mq *MongoQueries) AddHotList(ctx context.Context, collectionName string, hotList *hotListParams) (*HotList, error) {
	if hotList.HotlistTab == "" || hotList.URL == "" {
		return nil, errors.New("HotlistTab and URL must be provided")
	}

	// Check for duplicates
	filter := bson.M{
		"$or": []bson.M{
			{"hotlistTab": hotList.HotlistTab},
			{"url": hotList.URL},
		},
	}

	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("hot search with the same name or URL already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document
	newHotSearch := &HotList{
		ID:         primitive.NewObjectID(),
		HotlistTab: hotList.HotlistTab,
		URL:        hotList.URL,
		Name:       hotList.Name,
		ImageLink:  hotList.ImageLink,
		InfoTexts:  hotList.InfoTexts,
		CreatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
		UpdatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
	}

	// Insert the document
	err = mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		result, err := collection.InsertOne(ctx, newHotSearch)
		if err == nil {
			newHotSearch.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert hot search: %v", err)
	}

	return newHotSearch, nil

}

func (mq *MongoQueries) AddManyHotList(ctx context.Context, collectionName string, hotList []*hotListParams) ([]*HotList, error) {
	if len(hotList) == 0 {
		return nil, errors.New("no hot searches provided to insert")
	}

	var docs []interface{}
	var resultDocs []*HotList

	for _, params := range hotList {
		// Validate fields
		if params.ImageLink == "" || params.URL == "" {
			return nil, errors.New("ImageLink and URL must be provided")
		}

		newHotSearch := &HotList{
			ID:         primitive.NewObjectID(),
			HotlistTab: params.HotlistTab,
			URL:        params.URL,
			Name:       params.Name,
			ImageLink:  params.ImageLink,
			InfoTexts:  params.InfoTexts,
			CreatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
			UpdatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
		}

		docs = append(docs, newHotSearch)
		resultDocs = append(resultDocs, newHotSearch)
	}

	// Insert the documents
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		_, err := collection.InsertMany(ctx, docs)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert hot searches: %v", err)
	}

	return resultDocs, nil
}
