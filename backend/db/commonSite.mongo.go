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

//same with siteItem

type commonSiteParams struct {
	Name string `bson:"name" json:"name"`
	Icon string `bson:"icon" json:"icon"`
	Url  string `bson:"url" json:"url"`
}

func (mq *MongoQueries) GetAllCommonSiteItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddCommonSiteItem(ctx context.Context, collectionName string, commonSiteItem *commonSiteParams) (*CommonSite, error) {
	// Validate input
	if commonSiteItem.Name == "" || commonSiteItem.Url == "" {
		return nil, errors.New("name  URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": commonSiteItem.Name},
			{"url": commonSiteItem.Url},
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

	newSiteItems := &CommonSite{
		ID:   primitive.NewObjectID(),
		Name: commonSiteItem.Name,
		Icon: commonSiteItem.Icon,
		Url:  commonSiteItem.Url,
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

func (mq *MongoQueries) AddManyCommonSiteItem(ctx context.Context, collectionName string, commonSiteItems []*commonSiteParams) ([]*CommonSite, error) {
	// Validate input
	if len(commonSiteItems) == 0 {
		return nil, errors.New("no common site Items provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*CommonSite

	for _, commonSiteItem := range commonSiteItems {
		// Validate fields
		if commonSiteItem.Name == "" || commonSiteItem.Url == "" {
			return nil, errors.New("name and  URL must be provided")
		}

		newSiteItem := &CommonSite{
			ID:   primitive.NewObjectID(),
			Name: commonSiteItem.Name,
			Icon: commonSiteItem.Icon,
			Url:  commonSiteItem.Url,
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
