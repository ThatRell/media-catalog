package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("env variable 'MONGODB_URI' is unset")
	}

	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetCompressors([]string{"zstd"})

	// Creates a new client and connects to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("error connecting to mongo: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("error disconnecting to mongo: %v", err)
		}
	}()

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatalf("error pinging mongo: %v", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// POST add new media
	// GET retrieve all media
	// GET retive media by id
	// PUT update media with id
	// DELETE delete media with id

	r.POST("/media", func(c *gin.Context) {

	})

	r.GET("/media", func(c *gin.Context) {

	})

	r.GET("/media/:id", func(c *gin.Context) {
		id := c.Param("id")
	})

	r.PUT("/media/:id", func(c *gin.Context) {
		id := c.Param("id")
	})

	r.DELETE("/media/:id", func(c *gin.Context) {
		id := c.Param("id")
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
