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

type feedTopNewsParams struct {
	Url       string `bson:"url" json:"url"`
	Name      string `bson:"name" json:"name"`
	FeedTitle string `bson:"feedTitle" json:"feedTitle"`
}

func (mq *MongoQueries) GetAllTopNews(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddTopNews(ctx context.Context, collectionName string, feedTopNew *feedTopNewsParams) (*FeedTopNew, error) {
	if feedTopNew.Url == "" || feedTopNew.Name == "" || feedTopNew.FeedTitle == "" {
		return nil, errors.New("name, URL and feedTitle must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": feedTopNew.Name},
			{"url": feedTopNew.Url},
			{"feedTitle": feedTopNew.FeedTitle},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {

		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("failed to check unique constraint: %v", err)
		}
		if count > 0 {
			return errors.New("feed Title with the same name or URL already exists")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Prepare the new document

	newFeedTitle := &FeedTopNew{
		ID:        primitive.NewObjectID(),
		Name:      feedTopNew.Name,
		FeedTitle: feedTopNew.FeedTitle,
		Url:       feedTopNew.Url,
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
		result, err := collection.InsertOne(ctx, newFeedTitle)
		if err == nil {
			newFeedTitle.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert site items: %v", err)
	}

	return newFeedTitle, nil
}

func (mq *MongoQueries) AddManyTopNews(ctx context.Context, collectionName string, feedTopNew []*feedTopNewsParams) ([]*FeedTopNew, error) {
	// Validate input
	if len(feedTopNew) == 0 {
		return nil, errors.New("no common site Items provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*FeedTopNew

	for _, topNews := range feedTopNew {
		// Validate fields
		if topNews.Name == "" || topNews.Url == "" || topNews.FeedTitle == "" {
			return nil, errors.New("name, URL and feedTitle must be provided")
		}

		newFeedItem := &FeedTopNew{
			ID:        primitive.NewObjectID(),
			Name:      topNews.Name,
			FeedTitle: topNews.FeedTitle,
			Url:       topNews.Url,
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
		return nil, fmt.Errorf("failed to insert Feed Title: %v", err)
	}

	return resultDocs, nil
}
