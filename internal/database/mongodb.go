package database

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/yigit433/go-rest-template/internal/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *MongoAdapter

type MongoAdapter struct {
	Client *mongo.Client
	DbName string

	TasksCollection *mongo.Collection
}

func NewMongoAdapter() {
	dbUri := os.Getenv("MONGO_URI")
	if strings.TrimSpace(dbUri) == "" {
		logging.CustomLogger.Error("⚠️ 'MONGO_URI' is not set in the environment variables", map[string]interface{}{"eventId": "dbconfig_load"})
		return
	}

	dbName := os.Getenv("MONGO_DB")
	if strings.TrimSpace(dbName) == "" {
		logging.CustomLogger.Error("⚠️ 'MONGO_DB' is not set in the environment variables", map[string]interface{}{"eventId": "dbconfig_load"})
		return
	}

	clientCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientCtx, options.Client().ApplyURI(dbUri))
	if err != nil {
		logging.CustomLogger.Error("⚠️ Database connection failed! Check settings and permissions", map[string]interface{}{"eventId": "connect_db"})
		return
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if pingErr := client.Ping(pingCtx, nil); pingErr != nil {
		logging.CustomLogger.Error("⚠️ Database ping failed! Check connection and configuration", map[string]interface{}{"eventId": "connect_db"})
		return
	}

	logging.CustomLogger.Info("✅ Database connection established successfully!", map[string]interface{}{"eventId": "connect_db", "dbName": dbName, "provider": "MongoDB"})

	Mongo = &MongoAdapter{
		Client: client,
		DbName: dbName,

		TasksCollection: client.Database(dbName).Collection("tasks"),
	}
}

func GetCollection(dbClient *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return dbClient.Database(dbName).Collection(collectionName)
}