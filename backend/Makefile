DATABASE_URL=mongodb://admin:password@127.0.0.1:27017/?authSource=admin

delete_collection:
	echo "db.getSiblingDB('$(DB_NAME)').$(COLLECTION_NAME).drop();" | mongosh "$(DATABASE_URL)"

go_test:
	go test ./db -v

.PHONY:delete_collection go_test