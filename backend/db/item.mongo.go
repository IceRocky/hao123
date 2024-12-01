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

type ItemParams struct {
	Name     string  `bson:"name" json:"name"`
	URL      string  `bson:"url" json:"url"`
	Category *string `bson:"category,omitempty" json:"category,omitempty"`
}

func (mq *MongoQueries) GetAllItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
	// Default filter to empty (matches all documents) if no filter is provided
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

func (mq *MongoQueries) AddItem(ctx context.Context, collectionName string, item *ItemParams) (*Item, error) {
	// Validate input
	if item.Name == "" || item.URL == "" {
		return nil, errors.New("name and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": item.Name},
			{"url": item.URL},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {

		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("item with the same name or URL already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document

	newItemCategories := &Item{
		ID:       primitive.NewObjectID(),
		Name:     item.Name,
		URL:      item.URL,
		Category: item.Category,
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
		result, err := collection.InsertOne(ctx, newItemCategories)
		if err == nil {
			newItemCategories.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}

	return newItemCategories, nil

}

func (mq *MongoQueries) AddManyItem(ctx context.Context, collectionName string, items []*ItemParams) ([]*Item, error) {
	// Validate input
	if len(items) == 0 {
		return nil, errors.New("no items provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*Item

	for _, item := range items {
		// Validate fields

		if item.Name == "" || item.URL == "" {
			return nil, errors.New("all items must have a name and URL")
		}

		newItem := &Item{
			ID:   primitive.NewObjectID(),
			Name: item.Name,
			URL:  item.URL,
			CreatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
			UpdatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
		}

		docs = append(docs, newItem)
		resultDocs = append(resultDocs, newItem)
	}

	// Insert the documents
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		_, err := collection.InsertMany(ctx, docs)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %v", err)
	}

	return resultDocs, nil
}
