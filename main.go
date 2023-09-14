package main

import (
	"log"

	"api-go-lang-transations/controllers"
	seeds "api-go-lang-transations/database/seed"
	"api-go-lang-transations/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)

	// Create a new instance of the TransactionSeed struct
	transactionSeed := &seeds.TransactionSeed{}

	// Call the Seed method to add the transaction seed to the database
	transactionSeed.Seed(initializers.DB)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/transactions", func(router fiber.Router) {
		router.Post("/", controllers.CreateTransaction)
		router.Get("", controllers.ListTransactions)
	})
	micro.Route("/transactions/:transactionsId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteTransaction)
		router.Get("", controllers.FindTransactionById)
		router.Patch("", controllers.UpdateTransaction)
	})
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
