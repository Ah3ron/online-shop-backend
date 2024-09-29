package main

import (
	"online-shop/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db *pg.DB

func initDB() {
	db = pg.Connect(&pg.Options{
		Addr:     "autorack.proxy.rlwy.net:37883",
		User:     "postgres",
		Password: "rrzOKOQzFzQWWcLCyBiemZlFpXGehfug",
		Database: "railway",
	})

	// Убедитесь, что таблица создается только один раз
	err := db.Model(&models.Product{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true, // Создавать таблицу только если она не существует
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	initDB()
	defer db.Close()

	app.Post("/products", func(c *fiber.Ctx) error {
		product := new(models.Product)
		if err := c.BodyParser(product); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		_, err := db.Model(product).Insert()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot insert product"})
		}

		return c.Status(fiber.StatusCreated).JSON(product)
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		var products []models.Product
		err := db.Model(&products).Select()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot retrieve products"})
		}

		return c.JSON(products)
	})

	app.Listen(":3000")
}
