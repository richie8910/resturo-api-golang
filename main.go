package main

import (
	"os"

	"github.com/gin-gonic/gin"
	/* "github.com/sanketghosh/resturo-go-api/controllers"
	"github.com/sanketghosh/resturo-go-api/database" */)

func main() {

	// PORT declaration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

}
