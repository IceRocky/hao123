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

type topListItemParams struct {
	Url         string `bson:"url" json:"url"`
	Name        string `bson:"name" json:"name"`
	TopListName string `bson:"topListName" json:"topListName"`
}

func (mq *MongoQueries) GetAllTopListItem(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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
func (mq *MongoQueries) AddTopListItem(ctx context.Context, collectionName string, topListItem *topListItemParams) (*TopListItem, error) {
	if topListItem.Url == "" || topListItem.TopListName == "" {
		return nil, errors.New("name, URL and TopListName must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": topListItem.Name},
			{"url": topListItem.Url},
			{"topListName": topListItem.TopListName},
		},
	}

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

	newTopListItem := &TopListItem{
		ID:          primitive.NewObjectID(),
		Name:        topListItem.Name,
		TopListName: topListItem.TopListName,
		Url:         topListItem.Url,
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
		result, err := collection.InsertOne(ctx, newTopListItem)
		if err == nil {
			newTopListItem.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert site items: %v", err)
	}

	return newTopListItem, nil
}
func (mq *MongoQueries) AddManyTopListItem(ctx context.Context, collectionName string, topListItems []*topListItemParams) ([]*TopListItem, error) {
	if len(topListItems) == 0 {
		return nil, errors.New("no common site Items provided to insert")
	}

	// Prepare documents for insertion

	var docs []interface{}
	var resultDocs []*TopListItem

	for _, topList := range topListItems {
		// Validate fields
		if topList.Url == "" || topList.TopListName == "" {
			return nil, errors.New("name, URL and TopListName must be provided")
		}

		newTopList := &TopListItem{
			ID:          primitive.NewObjectID(),
			Name:        topList.Name,
			TopListName: topList.TopListName,
			Url:         topList.Url,
			CreatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
			UpdatedAt: func() *time.Time {
				now := time.Now()
				return &now
			}(),
		}

		docs = append(docs, newTopList)
		resultDocs = append(resultDocs, newTopList)
	}
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		_, err := collection.InsertMany(ctx, docs)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert Top List Items: %v", err)
	}

	return resultDocs, nil
}
