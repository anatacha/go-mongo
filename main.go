package main

import (
	"context"

	// "fmt"
	// "log"

	// "fmt"
	"go_mon/api"
	"go_mon/database"
	m "go_mon/model"

	// "go_mon/setting"

	// "github.com/patcharp/golib/requests"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

// 1).ใช้go + mongoเข้าสู่ขั้นตอนการเขียนโค๊ดเพื่อเชื่อมต่อกับ mongodb
// 2).สร้าง collection ยังไง

func createCollectionWithSchema() error {
	db := database.DB
	// ==CreateMany==
	// 1x สร้าง arr struct เก็บค่า ชื่อcollection กับ schema
	collections := []struct {
		Name   string
		Schema interface{}
	}{
		{Name: "movie", Schema: m.Movie{}},
		{Name: "music", Schema: m.Music{}},
		{Name: "game", Schema: m.Game{}},
	}
	// 2x multi create collection
	for _, coll := range collections {
		// 3x วน for เอาชื่อ collection มา สร้าง
		collection := db.Collection(coll.Name)
		// 4x ให้ ไล่ตามลำดับ _id 1น้อย->มาก -1มาก->น้อย ใช้ method IndexModel ของฟิวส์_id
		indexModel := mongo.IndexModel{
			Keys: bson.M{"_id": 1},
		}
		// 5x ใช้ในการสร้างหลาย Index พร้อมกัน ให้การสร้าง collection หลายๆอันให้ index _id แต่ละ collection มันเรียงจาก น->ม
		indexModels := []mongo.IndexModel{indexModel}
		_, err := collection.Indexes().CreateMany(context.Background(), indexModels)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	// ==CreateOne==
	collectionName := "book"
	collectionBook := db.Collection(collectionName)
	// 1x จัดลำดับ ของฟิวส์ _id
	indexModel := mongo.IndexModel{
		Keys: bson.M{"_id": 1},
	}

	_, err := collectionBook.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		logrus.Fatal(err)
	}
	return nil
}

func main() {
	// 1x Connect Db
	database.Connect()
	// 2x สร้าง Collection names and schemas
	createCollectionWithSchema()
	// 3x cmd => load env data,connect mongo,migrate collection,api
	cmdStartServer()
}

func cmdStartServer() error {
	// Connect
	database.Connect()
	// Migrate
	createCollectionWithSchema()
	// Api
	api.Register()
	return nil
}
