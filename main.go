package main

import (
	"books-management/config"
	"books-management/kafka"
	"books-management/migrations"
	"books-management/routes"
	"log"
	"runtime/debug"

	_ "books-management/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	debug.SetGCPercent(-1) // Disable optimizations

	loadEnvVariables()
	startKafkaConsumer()
	initializeKafka()
	initializeDatabase()
	runMigrations()
	startServer()
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func startKafkaConsumer() {
	go func() {
		log.Println("Starting Kafka Consumer...")
		kafka.Consume("book_events")
	}()
}

func initializeKafka() {
	kafka.InitKafka()
}

func initializeDatabase() {
	config.InitConfig()
}

func runMigrations() {
	migrations.RunMigrations()
}

func startServer() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Books API is running!"})
	})

	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
