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

type AddCategoryParams struct {
	Name string `bson:"name" json:"name"` // Name of the category
	URL  string `bson:"url" json:"url"`   // URL for the category
}

func (mq *MongoQueries) GetAllItemCategories(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddItemCategories(ctx context.Context, collectionName string, itemCategories *AddCategoryParams) (*ItemCategories, error) {
	// Validate input
	if itemCategories.Name == "" || itemCategories.URL == "" {
		return nil, errors.New("name and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": itemCategories.Name},
			{"url": itemCategories.URL},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("category with the same name or URL already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document

	newItemCategories := &ItemCategories{
		ID:   primitive.NewObjectID(),
		Name: itemCategories.Name,
		URL:  itemCategories.URL,
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
		return nil, fmt.Errorf("failed to insert item category: %v", err)
	}

	return newItemCategories, nil

}

func (mq *MongoQueries) AddManyItemCategories(ctx context.Context, collectionName string, itemCategories []*AddCategoryParams) ([]*ItemCategories, error) {
	// Validate input
	if len(itemCategories) == 0 {
		return nil, errors.New("no categories provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*ItemCategories

	for _, category := range itemCategories {
		// Validate fields
		if category.Name == "" || category.URL == "" {
			return nil, errors.New("all categories must have a name and URL")
		}

		newCategory := &ItemCategories{
			ID:   primitive.NewObjectID(),
			Name: category.Name,
			URL:  category.URL,
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
		return nil, fmt.Errorf("failed to insert item categories: %v", err)
	}

	return resultDocs, nil
}
