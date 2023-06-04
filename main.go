package main

import (
	"context"

	// "fmt"
	// "log"

	// "fmt"
	"go_mon/database"
	m "go_mon/model"

	// "github.com/patcharp/golib/requests"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/gofiber/fiber/v2"
	"github.com/patcharp/golib/v2/helper"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Age  int                `json:"age" bson:"age"`
}

// 1).ใช้go + mongoเข้าสู่ขั้นตอนการเขียนโค๊ดเพื่อเชื่อมต่อกับ mongodb
// 2).สร้าง collection ยังไง
func insertDataEP(c *fiber.Ctx) error {

	var persons Person

	if err := c.BodyParser(&persons); err != nil {
		return err
	}
	personsCollection := database.DB.Collection("persons")
	_, err := personsCollection.InsertOne(context.Background(), persons)
	if err != nil {
		return err
	}

	return helper.HttpOk(c, persons)

}

func getPersonEP(c *fiber.Ctx) error {

	collection := database.DB.Collection("persons")
	// bson.M{"name": "data1"}
	dataAll, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving users"})
	}
	var persons []Person
	if err := dataAll.All(context.Background(), &persons); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving persons"})
	}

	return c.JSON(persons)
	// return helper.HttpOk(c, result)

}

func updateUserEP(c *fiber.Ctx) error {

	id := c.Params("id")
	collection := database.DB.Collection("persons")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Object ID",
		})
	}

	var persons Person
	if err := c.BodyParser(&persons); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	updatePerson := bson.M{
		"$set": bson.M{
			"name": persons.Name,
			"age":  persons.Age,
		},
	}

	filter := bson.M{"_id": objectID}

	result, err := collection.UpdateOne(context.Background(), filter, updatePerson)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update persons"})
	}

	return c.JSON(result)
}
func checkCollectionExists(db *mongo.Database, collectionName string) (bool, error) {
	filter := bson.M{"book": collectionName}
	cursor, err := db.ListCollections(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return cursor.Next(context.Background()), nil
}

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
func getUserByIdEP(c *fiber.Ctx) error {
	id := c.Params("id")
	collection := database.DB.Collection("persons")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Object ID",
		})
	}

	// Define the filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Retrieve the user from MongoDB
	result := collection.FindOne(context.Background(), filter)
	var persons Person
	err = result.Decode(&persons)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(persons)
}
func main() {
	// 1x Connect Db
	database.Connect()
	// 2x สร้าง Collection names and schemas
	createCollectionWithSchema()

	// Set up your GoFiber application
	app := fiber.New()
	api := app.Group("/api")
	version1 := api.Group("/v1")
	// TODO ย้ายฟยร
	version1.Post("/data", insertDataEP)
	version1.Get("/data", getPersonEP)
	version1.Put("/data/:id", updateUserEP)
	version1.Get("/data/:id", getUserByIdEP)
	// Start the GoFiber server
	logrus.Fatal(app.Listen(":3000"))
}
