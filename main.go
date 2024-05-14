package main

import (
	// "github.com/Qwerci/eos-api2/models"
	"github.com/Qwerci/eos-api2/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.Default())


	router.POST("/createfield", controllers.CreateField)


	router.Run(":3033")
}