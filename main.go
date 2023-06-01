package main

import (
	"context"
	"fmt"

	// "github.com/patcharp/golib/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/patcharp/golib/v2/helper"
	"github.com/sirupsen/logrus"

	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
type Music struct {
	gorm.Model
	Name string
}

func insertData(c *fiber.Ctx) error {
	// เข้าสู่ขั้นตอนการเขียนโค๊ดเพื่อเชื่อมต่อกับ mongodb
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	var result = map[string]interface{}{}
	if err != nil {
		logrus.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Close the MongoDB client when the program exits
	defer client.Disconnect(context.TODO())

	fmt.Println("Document inserted successfully!")

	persons := new(Person)
	if err := c.BodyParser(persons); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user data"})
	}
	//DB Collection
	personsCollection := client.Database("golang-test").Collection("persons")
	_, err = personsCollection.InsertOne(context.Background(), persons)
	if err != nil {
		return err
	}
	result = map[string]interface{}{
		"name": persons,
	}
	return helper.HttpOk(c, result)
}

func main() {

	// Set up your GoFiber application
	app := fiber.New()
	app.Post("/data", insertData)

	// Start the GoFiber server
	logrus.Fatal(app.Listen(":3000"))
}
