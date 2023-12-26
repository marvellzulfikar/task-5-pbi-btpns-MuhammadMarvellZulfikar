package main

import (
	"github.com/marvellzulfikar/task-5-pbi-btpns-MuhammadMarvellZulfikar/controllers/photoscontroller"
	"github.com/marvellzulfikar/task-5-pbi-btpns-MuhammadMarvellZulfikar/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/photos", photoscontroller.Index)
	r.POST("/api/photos", photoscontroller.Create)
	r.PUT("/api/photos/:id", photoscontroller.Update)
	r.DELETE("/api/photos", photoscontroller.Delete)

	r.Run()

}
