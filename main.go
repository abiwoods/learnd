package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var db dataSource

func main() {
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatalf("cannot start DB: %w", err)
	}

	router := gin.Default()
	setHandlers(router)

	// In production would make this configurable,
	// for this instance, feels sufficient to use default value of 8080
	if err := router.Run(); err != nil {
		log.Fatalf("cannot start server: %w", err)
	}
}

func setHandlers(router *gin.Engine) {
	router.GET("/customers/:customerID/meters", getMetersHandler)
	router.GET("/customers/:customerID/meters/:serialID/usage", getUsageHandler)
}
