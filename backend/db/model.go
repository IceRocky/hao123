package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SiteItems struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Icon      string             `bson:"icon" json:"icon"`
	Link      string             `bson:"link" json:"link"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type Item struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	URL       string             `bson:"url" json:"url"`                               // URL or link
	Icon      *string            `bson:"icon,omitempty" json:"icon,omitempty"`         // Optional icon for the item
	Category  *string            `bson:"category,omitempty" json:"category,omitempty"` // Optional category name
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type HostSearch struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Link      string             `bson:"link" json:"link"` // The URL
	Title     string             `bson:"title" json:"title"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"` // Unique identifier for the category
	Name      string             `bson:"name" json:"name"`        // Name of the category
	URL       string             `bson:"url" json:"url"`          // URL for the category
	Items     []Item             `bson:"items" json:"items"`      // List of items under the category
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type ItemCategories struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"` // Name of the category
	URL       string             `bson:"url" json:"url"`   // URL for the category
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type CommonSite struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Icon      string             `bson:"icon" json:"icon"`
	Url       string             `bson:"url" json:"url"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type FeedTopNew struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Url       string             `bson:"url" json:"url"`
	Name      string             `bson:"name" json:"name"`
	FeedTitle string             `bson:"feedTitle" json:"feedTitle"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type GovSites struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Url       string             `bson:"url" json:"url"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type HotList struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	HotlistTab string             `bson:"hotlistTab" json:"hotlistTab"`
	URL        string             `bson:"url" json:"url"`
	Name       string             `bson:"name" json:"name"`
	ImageLink  string             `bson:"imageLink" json:"imageLink"`
	InfoTexts  []string           `bson:"infoTexts" json:"infoTexts"`
	CreatedAt  *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type HotListTab struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Url       string             `bson:"url" json:"url"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type TopList struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Url       string             `bson:"url" json:"url"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

type TopListItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TopListName string             `bson:"topListName" json:"topListName"`
	Url         string             `bson:"url" json:"url"`
	Name        string             `bson:"name" json:"name"`
	CreatedAt   *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt   *time.Time         `bson:"updatedAt" json:"updatedAt"`
}
