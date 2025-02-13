package database

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env, %v", err)
	}

	MONGO_DB := os.Getenv("MONGOdb_URL")

	
}
