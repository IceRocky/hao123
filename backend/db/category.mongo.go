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

//	type CategoryItemParams struct {
//		Name     string  `bson:"name" json:"name"`
//		URL      string  `bson:"url" json:"url"`
//		Category *string `bson:"category" json:"category"`
//	}
//
//	type CategoryParams struct {
//		Name  string               `bson:"name" json:"name"` // Name of the category
//		URL   string               `bson:"url" json:"url"`
//		Items []CategoryItemParams `bson:"categoryitem" json:"item"`
//	}
type CategoryItemParams struct {
	Name     string  `bson:"name" json:"name"`
	URL      string  `bson:"url" json:"url"`
	Category *string `bson:"category" json:"category"` // URL or link
}

// CategoryParams represents the structure of a category with a list of items
type CategoryParams struct {
	Name  string               `bson:"name" json:"name"`   // Name of the category
	URL   string               `bson:"url" json:"url"`     // URL for the category
	Items []CategoryItemParams `bson:"items" json:"items"` // List of items under the category

}

func (mq *MongoQueries) GetAllCategories(ctx context.Context, collectionName string, filter bson.M, skip, limit int64) ([]bson.M, error) {
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

func (mq *MongoQueries) AddCategory(ctx context.Context, collectionName string, category *CategoryParams) (*Category, error) {
	// Validate input
	if category.Name == "" || category.URL == "" {
		return nil, errors.New("name and URL must be provided")
	}

	// Define a filter to check for existing documents
	filter := bson.M{
		"$or": []bson.M{
			{"name": category.Name},
			{"url": category.URL},
		},
	}

	// Check for duplicates
	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		fmt.Println("Checking duplicates for:", collectionName)
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

	// Prepare the new category document

	var items []Item
	if category.Items != nil {

		for _, item := range category.Items {
			items = append(items, Item{
				Name:     item.Name,
				URL:      item.URL,
				Category: item.Category,
			})
		}

	}
	newCategory := &Category{
		ID:    primitive.NewObjectID(),
		Name:  category.Name,
		URL:   category.URL,
		Items: items,
		CreatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
		UpdatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
	}

	// Insert the category document
	err = mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		result, err := collection.InsertOne(ctx, newCategory)
		if err == nil {
			newCategory.ID = result.InsertedID.(primitive.ObjectID)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert category: %v", err)
	}

	return newCategory, nil
}

func (mq *MongoQueries) AddManyCategories(ctx context.Context, collectionName string, categories []*CategoryParams) ([]*Category, error) {
	// Validate input
	if len(categories) == 0 {
		return nil, errors.New("no categories provided to insert")
	}

	// Prepare documents for insertion
	var docs []interface{}
	var resultDocs []*Category

	err := mq.ExecuteQuery(ctx, collectionName, func(collection *mongo.Collection) error {
		for _, category := range categories {
			// Validate fields
			if category.Name == "" || category.URL == "" {
				return errors.New("all categories must have a name and URL")
			}

			// Check if the category already exists
			filter := bson.M{"name": category.Name, "url": category.URL}
			count, err := collection.CountDocuments(ctx, filter)
			if err != nil {
				return fmt.Errorf("failed to check for existing category: %v", err)
			}
			if count > 0 {
				fmt.Printf("Category '%s' already exists, skipping...\n", category.Name)
				continue
			}

			// Prepare items
			var items []Item
			if category.Items != nil {
				for _, item := range category.Items {
					items = append(items, Item{
						Name:     item.Name,
						URL:      item.URL,
						Category: item.Category,
					})
				}
			}

			// Create new Category object
			newCategory := &Category{
				ID:    primitive.NewObjectID(),
				Name:  category.Name,
				URL:   category.URL,
				Items: items,
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

		// Insert only if there are new documents
		if len(docs) > 0 {
			_, err := collection.InsertMany(ctx, docs)
			if err != nil {
				return fmt.Errorf("failed to insert categories: %v", err)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resultDocs, nil
}

//func TestAddManyCategory(t *testing.T) {
//	category, err := LoadFromFile[CategoryParams]("../sample-data/category.json")
//
//	if err != nil {
//		panic(fmt.Sprintf("Failed to load categories: %v", err))
//	}
//
//	var dummy []*CategoryParams
//
//	for i := range category {
//		dummy = append(dummy, &category[i])
//	}
//
//	insertCategory, err := testStore.AddManyCategories(context.Background(), "categories", dummy)
//	require.NoError(t, err) // Ensure no errors occurred
//	require.NotEmpty(t, insertCategory)
//
//	for i, inserted := range insertCategory {
//		require.NotNil(t, inserted.ID)
//		require.Equal(t, dummy[i].Name, inserted.Name)
//		require.Equal(t, dummy[i].URL, inserted.URL)
//		require.NotNil(t, inserted.CreatedAt)
//		require.NotNil(t, inserted.UpdatedAt)
//
//		// Validate items within each category, if present
//		if len(dummy[i].Items) > 0 {
//			require.Len(t, inserted.Items, len(dummy[i].Items))
//			for j, item := range inserted.Items {
//				require.Equal(t, dummy[i].Items[j].Name, item.Name)
//				require.Equal(t, dummy[i].Items[j].URL, item.URL)
//				require.Equal(t, *dummy[i].Items[j].Category, *item.Category)
//			}
//		}
//	}
//}
