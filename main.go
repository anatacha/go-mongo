package main

import (
	"context"
	// "fmt"
	"go_mon/database"
	m "go_mon/model"

	// "github.com/patcharp/golib/requests"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/gofiber/fiber/v2"
	"github.com/patcharp/golib/v2/helper"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	// `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

// 1).ใช้go + mongoเข้าสู่ขั้นตอนการเขียนโค๊ดเพื่อเชื่อมต่อกับ mongodb
// 2).สร้าง collection ยังไง
func insertData(c *fiber.Ctx) error {

	var persons Person

	var results map[string]interface{}
	if err := c.BodyParser(&persons); err != nil {
		return err
	}
	personsCollection := database.DB.Collection("persons")
	_, err := personsCollection.InsertOne(context.Background(), persons)
	if err != nil {
		return err
	}
	results = map[string]interface{}{
		"name": persons,
	}
	return helper.HttpOk(c, results)

}

func getPersonEp(c *fiber.Ctx) error {

	collection := database.DB.Collection("persons")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving users"})
	}
	defer cursor.Close(context.Background())

	var persons []Person
	if err := cursor.All(context.Background(), &persons); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving persons"})
	}

	return c.JSON(persons)
	// return helper.HttpOk(c, result)

}

func createCollectionWithSchema() error {

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
		collection := database.DB.Collection(coll.Name)
		// 4x ให้ ไล่ตามลำดับ _id 1น้อย->มาก -1มาก->น้อย ใช้ method IndexModel
		indexModel := mongo.IndexModel{
			Keys: bson.M{"_id": 1},
		}
		// 5x arr [{"_id": 1}]
		indexModels := []mongo.IndexModel{indexModel}
		_, err := collection.Indexes().CreateMany(context.Background(), indexModels)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	// one
	// Create the collection
	// collectionName := "book"
	// err := database.DB.CreateCollection(context.Background(), collectionName)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	return nil
}
func main() {
	// 1x Connect Db
	database.Connect()
	// 2x สร้าง Collection names and schemas
	createCollectionWithSchema()

	// Set up your GoFiber application
	app := fiber.New()
	app.Post("/data", insertData)
	app.Get("/data", getPersonEp)

	// Start the GoFiber server
	logrus.Fatal(app.Listen(":3000"))
}
