package v1

import (
	// "fmt"
	"context"
	"go_mon/database"
	m "go_mon/model"

	"github.com/gofiber/fiber/v2"
	"github.com/patcharp/golib/v2/helper"

	// "github.com/pdfcpu/pdfcpu/pkg/api"
	"os/exec"

	"github.com/go-pdf/fpdf"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertDataEP(c *fiber.Ctx) error {
	var persons m.Person

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

func GetPersonEP(c *fiber.Ctx) error {

	collection := database.DB.Collection("persons")
	// bson.M{"name": "data1"}
	dataAll, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving users"})
	}
	var persons []m.Person
	if err := dataAll.All(context.Background(), &persons); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving persons"})
	}

	return c.JSON(persons)
	// return helper.HttpOk(c, result)

}

func UpdateUserEP(c *fiber.Ctx) error {

	id := c.Params("id")
	collection := database.DB.Collection("persons")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Object ID",
		})
	}

	var persons m.Person
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

func UserByIdEP(c *fiber.Ctx) error {
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
	var persons m.Person
	err = result.Decode(&persons)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(persons)
}

func PdfEP(c *fiber.Ctx) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	errors := exec.Command("xdg-open", "hello.pdf").Start()
	if err != nil {
		logrus.Fatal(errors)
	}
	return err
}

func PdfFileEP(c *fiber.Ctx) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.ImageOptions("Add.pdf", 0, 0, 210, 297, false, fpdf.ImageOptions{}, 0, "")


	return nil
}
