package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func createGovSites(t *testing.T) *GovSites {
	arg := &govSitesParams{
		Url:  "http://tuijian.hao123.com/",
		Name: "hao123推荐",
	}

	result, err := testStore.AddGovSite(context.Background(), "govSites", arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.Name, result.Name)
	require.Equal(t, arg.Url, result.Url)

	return result
}

func TestAddGovSite(t *testing.T) {
	createGovSites(t)
}

func TestGetAllGovSites(t *testing.T) {
	insertGovSites := createGovSites(t)
	filter := bson.M{}
	results, err := testStore.GetAllGovSites(context.Background(), "govSites", filter, 0, 3)

	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Ensure at least one result matches the inserted category
	var found bool
	for _, result := range results {
		if result["url"] == insertGovSites.Url && result["name"] == insertGovSites.Name {
			found = true
			break
		}
	}

	require.True(t, found, "Inserted insertGovSites not found in the results")
}

func TestAddManyGovSites(t *testing.T) {
	govSites, err := LoadFromFile[govSitesParams]("../sample-data/gov-sites.json")

	if err != nil {
		panic(fmt.Sprintf("Failed to load categories: %v", err))
	}

	var dummy []*govSitesParams
	for i := range govSites {
		dummy = append(dummy, &govSites[i])
	}
	require.NotEmpty(t, dummy, "Dummy slice should not be empty")
	require.Equal(t, len(govSites), len(dummy), "Dummy slice should contain the same number of items as the loaded data")

	insertSiteItem, err := testStore.AddManyGovSites(context.Background(), "govSites", dummy)

	require.NoError(t, err)
	require.Len(t, insertSiteItem, len(dummy))
}
