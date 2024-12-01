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

type hotTabParams struct {
	Url  string `bson:"url" json:"url"`
	Name string `bson:"name" json:"name"`
}

func (mq *MongoQueries) GetAllHotTab(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddHotTab(ctx context.Context, collectionName string, hotTab *hotTabParams) (*HotListTab, error) {
	if hotTab.Url == "" || hotTab.Name == "" {
		return nil, errors.New("name and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": hotTab.Name},
			{"url": hotTab.Url},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {

		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("gov Title with the same name or URL already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document

	hotListsTab := &HotListTab{
		ID:   primitive.NewObjectID(),
		Name: hotTab.Name,
		Url:  hotTab.Url,
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
		result, err := collection.InsertOne(ctx, hotListsTab)
		if err == nil {
			hotListsTab.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert site items: %v", err)
	}

	return hotListsTab, nil
}

func (mq *MongoQueries) AddManyHotTab(ctx context.Context, collectionName string, hotTab []*hotTabParams) ([]*HotListTab, error) {
	// Validate input
	if len(hotTab) == 0 {
		return nil, errors.New("no common gov sites provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*HotListTab

	for _, tab := range hotTab {
		// Validate fields
		if tab.Name == "" || tab.Url == "" {
			return nil, errors.New("name and URL must be provided")
		}

		newFeedItem := &HotListTab{
			ID:   primitive.NewObjectID(),
			Name: tab.Name,
			Url:  tab.Url,
			CreatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
			UpdatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
		}

		docs = append(docs, newFeedItem)
		resultDocs = append(resultDocs, newFeedItem)
	}

	// Insert the documents
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		_, err := collection.InsertMany(ctx, docs)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert hot tab: %v", err)
	}

	return resultDocs, nil
}
