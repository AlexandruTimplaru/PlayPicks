package main

import (
	"log"

	"github.com/AlexandruTimplaru/PlayPicks/database"
	"github.com/AlexandruTimplaru/PlayPicks/health"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	g := gin.Default()
	api := g.Group("/api")

	health.Init(api)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
