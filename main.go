package main

import (
	muridController "gorm/controllers"
	"gorm/models"

	"github.com/gin-gonic/gin"
)

func main()  {
	Route := gin.Default()
	models.ConnectDatabase()
	Route.Static("/img", "./img")
	Route.MaxMultipartMemory = 8 << 20
	Route.GET("/api/murid", muridController.Index)
	Route.GET("/api/murid/:id", muridController.Show)
	Route.POST("/api/murid", muridController.Store)
	Route.PUT("/api/murid/:id", muridController.Update)
	Route.DELETE("/api/murid/:id", muridController.Destroy)

	Route.Run()
}