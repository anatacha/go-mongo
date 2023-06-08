package service

import (
	"context"
	"fmt"

	// "os"
	"go_mon/setting"

	"github.com/patcharp/golib/v2/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Database
	getEnv = util.GetEnv
)

func InitDb() error {
	// dbPort := getEnv("DB_PORT", "27017")
	// dbPort := getEnv("DB_PORT", "27017")
	// dbName := getEnv("DB_NAME", "golang-test")

	dbHost := setting.GetCfg().Db.Host
	// if err := setting.GetCfg().Load(); err != nil {
	// 	return err
	// }
	// dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := setting.GetCfg().Db.Port
	// dbName := setting.GetCfg().Db.Name
	dbName := getEnv("DB_NAME", "golang-test")

	// test := setting.GetCfg().DbHost
	// logrus.Infoln(test)
	// dbName := os.Getenv("DB_NAME")
	connectionURI := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	// สร้างหรือเรียกDB
	// DB = client.Database("golang-test")
	DB = client.Database(dbName)
	return nil
}
