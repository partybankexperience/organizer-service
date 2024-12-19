package test

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
)

//
//import (
//	"github.com/djfemz/rave/partybank-app/models"
//	"github.com/djfemz/rave/partybank-app/utils"
//	"log"
//	"testing"
//)
//
//func TestIsValidEndTime(t *testing.T) {
//	ticket := &models.Ticket{
//		ActivePeriod: &models.ActivePeriod{
//			EndDate:   "2024-03-01",
//			EndTime:   "09:00:00",
//			StartDate: "2024-01-24",
//			StartTime: "9:00",
//		},
//	}
//	status := utils.IsTicketSaleEndedFor(ticket)
//	log.Println(status)
//}


package main

import (
"errors"
"io"
"log"
"mime/multipart"

"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", UploadImage)
	router.GET("/image/:id", GetImage)
	if err := router.Run(":8080"); err != nil {
		log.Fatalln("error starting server: ", err)
	}
}

var images = make(map[string]string)

func UploadImage(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(400, errors.New("file not present in request"))
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("error: ", err)
		}
	}(file)

	bs, err := io.ReadAll(file)
	if err != nil {
		ctx.AbortWithStatusJSON(400, errors.New("could not read from file"))
		return
	}
	images["abc"] = string(bs)
	ctx.JSON(200, map[string]string{"url": "http://localhost:8080/image/abc"})
}

func GetImage(ctx *gin.Context) {
	imageId := ctx.Param("id")
	if imageId == "abc" {
		image := images["abc"]
		bs := []byte(image)
		ctx.Data(200, "image/png", bs)
		return
	}
	ctx.AbortWithStatusJSON(400, errors.New("not found"))
}

