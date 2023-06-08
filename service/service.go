package service

import (
	"context"
	"fmt"

	// "os"
	"go_mon/setting"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *mongo.Database
)

func InitDb() error {
	// dbHost := getEnv("DB_HOST", "127.0.0.1")
	// dbPort := getEnv("DB_PORT", "27017")
	// dbName := getEnv("DB_NAME", "golang-test")

	dbHost := setting.GetCfg().Db.Host
	dbPort := setting.GetCfg().Db.Port
	dbName := setting.GetCfg().Db.Name

	// connectionURI := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	connectMongo := constructMongoDBConnectionString(dbHost, dbPort)
	logrus.Infoln(connectMongo)
	clientOptions := options.Client().ApplyURI(connectMongo)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	// สร้างหรือเรียกDB
	// DB = client.Database("golang-test")
	DB = client.Database(dbName)
	return nil
}

// Function to construct MongoDB connection string
func constructMongoDBConnectionString(host, port string) string {
	return fmt.Sprintf("mongodb://%s:%s", host, port)

}
