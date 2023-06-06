package database

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := "27017"
	connectionURI := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	// สร้างหรือเรียกDB
	DB = client.Database("golang-test")
	return nil
}
