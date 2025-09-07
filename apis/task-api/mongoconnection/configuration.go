package mongoconnection

import "os"

// resolveMongoDbUri gets the MongoDB URI from environment or defaults to localhost
func resolveMongoDbUri() string {
	mongoUri := os.Getenv("MONGODB_URI")

	if mongoUri == "" {
		mongoUri = "mongodb://localhost:27017"
	}

	return mongoUri
}
