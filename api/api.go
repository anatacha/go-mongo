package api

import (
	v1 "go_mon/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Register() {
	app := fiber.New()
	api := app.Group("/api")
	version1 := api.Group("/v1")
	// DATA
	data := version1.Group("/data/")
	data.Get("", v1.GetPersonEP)
	data.Post("", v1.InsertDataEP)
	data.Put("/:id", v1.UpdateUserEP)
	data.Get("/:id", v1.UserByIdEP)
	logrus.Fatal(app.Listen(":3000"))
}
