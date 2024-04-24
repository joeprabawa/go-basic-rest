package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joeprabawa/basic-go-rest/database"
	model "github.com/joeprabawa/basic-go-rest/models"
)

func main() {

	database.Connect()
	// create new fiber app
	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1.Route("/destination", func(r fiber.Router) {
		r.Get("/", func(c *fiber.Ctx) error {
			destinations := []model.Destination{}
			all := database.Db.Find(&destinations)
			if all.RowsAffected == 0 {
				return c.Status(404).JSON("No destination found")
			}
			if all.Error != nil {
				return c.Status(500).JSON("Internal Server Error", all.Error.Error())
			}

			return c.Status(200).JSON(destinations)
		})
		r.Get("/:id", func(c *fiber.Ctx) error {
			params := c.Params("id")
			destinaton := []model.Destination{}
			find := database.Db.First(&destinaton, "id = ?", params)
			if find.RowsAffected == 0 {
				return c.Status(404).JSON("No destination found")
			}
			if find.Error != nil {
				return c.Status(500).JSON("Internal Server Error", find.Error.Error())
			}
			return c.Status(200).JSON(destinaton)
		})
		r.Post("/", func(c *fiber.Ctx) error {
			newDestination := new(model.Destination)
			parse := c.BodyParser(newDestination)
			if parse != nil {
				return c.Status(400).JSON(parse.Error())
			}

			create := database.Db.Create(&newDestination)

			if create.Error != nil {
				return c.Status(500).JSON("Internal Server Error", create.Error.Error())
			}

			return c.Status(201).JSON(newDestination)
		})
		r.Put("/:id", func(c *fiber.Ctx) error {
			params := c.Params("id")
			destination := new(model.Destination)
			parse := c.BodyParser(destination)
			if parse != nil {
				return c.Status(400).JSON(parse.Error())
			}
			put := database.Db.Model(&model.Destination{}).Where("id = ?", params).Updates(destination)
			if put.Error != nil {
				return c.Status(500).JSON("Internal Server Error", put.Error.Error())
			}
			return c.Status(200).JSON(destination)
		})
		r.Delete("/:id", func(c *fiber.Ctx) error {
			params := c.Params("id")
			destination := new(model.Destination)
			delete := database.Db.Where("id = ?", params).Delete(&destination)
			if delete.Error != nil {
				return c.Status(500).JSON("Internal Server Error", delete.Error.Error())
			}
			return c.Status(200).JSON(destination)

		})

	})

	app.Listen(":3000")
}
