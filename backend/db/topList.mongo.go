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

type topListParams struct {
	Url  string `bson:"url" json:"url"`
	Name string `bson:"name" json:"name"`
}

func (mq *MongoQueries) GetAllTopList(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddTopList(ctx context.Context, collectionName string, topList *topListParams) (*TopList, error) {
	if topList.Url == "" || topList.Name == "" {
		return nil, errors.New("name and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": topList.Name},
			{"url": topList.Url},
		},
	}
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

	list := &TopList{
		ID:   primitive.NewObjectID(),
		Name: topList.Name,
		Url:  topList.Url,
		CreatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
		UpdatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
	}

	err = mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		result, err := collection.InsertOne(ctx, list)
		if err == nil {
			list.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert site items: %v", err)
	}

	return list, nil

}

func (mq *MongoQueries) AddManyTopList(ctx context.Context, collectionName string, topList []*topListParams) ([]*TopList, error) {
	if len(topList) == 0 {
		return nil, errors.New("no common gov sites provided to insert")
	}

	var docs []interface{}
	var resultDocs []*TopList

	for _, top := range topList {
		// Validate fields
		if top.Name == "" || top.Url == "" {
			return nil, errors.New("name and URL must be provided")
		}

		newTopList := &TopList{
			ID:   primitive.NewObjectID(),
			Name: top.Name,
			Url:  top.Url,
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
		return nil, fmt.Errorf("failed to insert hot tab: %v", err)
	}

	return resultDocs, nil
}
