package db

import (
	"hao123/util"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../.")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	connect, err := Connect(config.DatabaseURL, config.DBName, config.DBUsername, config.DBPassword, config.Source, 3) // Use a test database
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func() {
		if err := connect.Disconnect(connect); err != nil {
			log.Printf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	queries := NewMongoQueries(connect.Database)       // Pass only the database
	testStore = NewMongoStore(connect.Client, queries) // Initialize test store

	code := m.Run()

	// Clean up the test database
	//err = connect.Database.Drop(context.Background())
	//if err != nil {
	//	log.Printf("Failed to clean up test database: %v", err)
	//}

	os.Exit(code)
}
