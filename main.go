package main

import (
	"log"

	"api-go-lang-transations/controllers"
	seeds "api-go-lang-transations/database/seed"
	"api-go-lang-transations/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "api-go-lang-transations/docs"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)

	transactionSeed := &seeds.TransactionSeed{}

	transactionSeed.Seed(initializers.DB)
}

// @title Swagger Transactions API
// @version 1.0
// @description This is a simple API write in Golang and Fiber to manage transactions used on Warren front-end test.

// @contact.name Johan Guse
// @contact.url https://johanguse.dev
// @contact.email johanguse@gmail.com

// @license MIT

// @host api-go-lang-transations.fly.dev
// @BasePath /api
// @schemes https
func main() {
	RUN_MODE := viper.GetString("Environment")
	if RUN_MODE == "" {
		RUN_MODE = "dev"
	}
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://api-go-lang-transations.fly.dev",
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
			"message": "Welcome to Golang, Fiber, and GORM. You are operating in " + RUN_MODE,
		})
	})

	micro.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Fatal(app.Listen(":8000"))
}
