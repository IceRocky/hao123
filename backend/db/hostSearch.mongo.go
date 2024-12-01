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

type HostSearchParams struct {
	Link  string `bson:"link" json:"link"`
	Title string `bson:"title" json:"title"`
}

func (mq *MongoQueries) GetAllHotSearch(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddHostSearch(ctx context.Context, collectionName string, hostSearch *HostSearchParams) (*HostSearch, error) {
	// Validate input
	if hostSearch.Link == "" || hostSearch.Title == "" {
		return nil, errors.New("link and title must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"link": hostSearch.Link},
			{"title": hostSearch.Title},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {

		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("host search with the same link or title already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document

	newItemCategories := &HostSearch{
		ID:    primitive.NewObjectID(),
		Link:  hostSearch.Link,
		Title: hostSearch.Title,
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
		return nil, fmt.Errorf("failed to insert host search: %v", err)
	}

	return newItemCategories, nil

}

func (mq *MongoQueries) AddManyHotSearch(ctx context.Context, collectionName string, hostSearch []*HostSearchParams) ([]*HostSearch, error) {
	// Validate input
	if len(hostSearch) == 0 {
		return nil, errors.New("no host search provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*HostSearch

	for _, search := range hostSearch {
		// Validate fields
		if search.Title == "" || search.Link == "" {
			return nil, errors.New("all host search must have a title and Link")
		}

		newCategory := &HostSearch{
			ID:    primitive.NewObjectID(),
			Link:  search.Link,
			Title: search.Title,

			CreatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
			UpdatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
		}

		docs = append(docs, newCategory)
		resultDocs = append(resultDocs, newCategory)
	}

	// Insert the documents
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		_, err := collection.InsertMany(ctx, docs)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert item hostSearch: %v", err)
	}

	return resultDocs, nil
}
