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

type govSitesParams struct {
	Url  string `bson:"url" json:"url"`
	Name string `bson:"name" json:"name"`
}

func (mq *MongoQueries) GetAllGovSites(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddGovSite(ctx context.Context, collectionName string, govSites *govSitesParams) (*GovSites, error) {
	if govSites.Url == "" || govSites.Name == "" {
		return nil, errors.New("name and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": govSites.Name},
			{"url": govSites.Url},
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

	newGovSites := &GovSites{
		ID:   primitive.NewObjectID(),
		Name: govSites.Name,
		Url:  govSites.Url,
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
		result, err := collection.InsertOne(ctx, newGovSites)
		if err == nil {
			newGovSites.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert site items: %v", err)
	}

	return newGovSites, nil
}

func (mq *MongoQueries) AddManyGovSites(ctx context.Context, collectionName string, govSites []*govSitesParams) ([]*GovSites, error) {
	// Validate input
	if len(govSites) == 0 {
		return nil, errors.New("no common gov sites provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*GovSites

	for _, govSite := range govSites {
		// Validate fields
		if govSite.Name == "" || govSite.Url == "" {
			return nil, errors.New("name and URL must be provided")
		}

		newFeedItem := &GovSites{
			ID:   primitive.NewObjectID(),
			Name: govSite.Name,
			Url:  govSite.Url,
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
		return nil, fmt.Errorf("failed to insert govSites: %v", err)
	}

	return resultDocs, nil
}
