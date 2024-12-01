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

type siteItemsParams struct {
	Name string `bson:"name" json:"name"`
	Icon string `bson:"icon" json:"icon"`
	Link string `bson:"link" json:"link"`
}

func (mq *MongoQueries) GetAllSiteItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddSiteItem(ctx context.Context, collectionName string, siteItem *siteItemsParams) (*SiteItems, error) {
	// Validate input
	if siteItem.Name == "" || siteItem.Icon == "" || siteItem.Link == "" {
		return nil, errors.New("name ,siteItem and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": siteItem.Name},
			{"icon": siteItem.Icon},
			{"link": siteItem.Link},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {

		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("siteItem with the same name or URL already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document

	newSiteItems := &SiteItems{
		ID:   primitive.NewObjectID(),
		Name: siteItem.Name,
		Icon: siteItem.Icon,
		Link: siteItem.Link,
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
		result, err := collection.InsertOne(ctx, newSiteItems)
		if err == nil {
			newSiteItems.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert site items: %v", err)
	}

	return newSiteItems, nil

}

func (mq *MongoQueries) AddManySiteItem(ctx context.Context, collectionName string, siteItems []*siteItemsParams) ([]*SiteItems, error) {
	// Validate input
	if len(siteItems) == 0 {
		return nil, errors.New("no site Items provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*SiteItems

	for _, siteItem := range siteItems {
		// Validate fields
		if siteItem.Name == "" || siteItem.Icon == "" || siteItem.Link == "" {
			return nil, errors.New("name ,siteItem and URL must be provided")
		}

		newSiteItem := &SiteItems{
			ID:   primitive.NewObjectID(),
			Name: siteItem.Name,
			Icon: siteItem.Icon,
			Link: siteItem.Link,
			CreatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
			UpdatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
		}

		docs = append(docs, newSiteItem)
		resultDocs = append(resultDocs, newSiteItem)
	}

	// Insert the documents
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		_, err := collection.InsertMany(ctx, docs)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert site Items: %v", err)
	}

	return resultDocs, nil
}
