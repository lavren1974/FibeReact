package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	mongodb := os.Getenv("MONGODB_URI")

	if port == "" {
		port = "3000"
	}

	if mongodb == "" {
		mongodb = "mongodb://localhost:27017/fibereact"
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb))
	coll := client.Database("fibereact").Collection("users")
	coll.InsertOne(context.TODO(), bson.D{{
		"Nik",
		"Lavr",
	}})

	if err != nil {
		panic(err)
	}

	app.Static("/", "./frontend/dist")

	app.Get("/app", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"app": "FibeReact",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"app": "FibeReact",
		})
	})

	log.Fatal(app.Listen(":" + port))
}
