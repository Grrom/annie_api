package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	saucenao "github.com/Grrom/annie_api/utils"
)

func getSauce(c *gin.Context) {
	saucenaoClient := saucenao.New()
	result, err := saucenaoClient.FromURL(c.Query("image_url"))

	if err != nil {
		log.Fatalf("Error: .Err:%s", err)
		return
	}

	similarity, err := strconv.ParseFloat(result.Data[0].Header.Similarity, 64)
	if err != nil {
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		saucenao.Sauce{
			Link:      result.Data[0].Data.ExtUrls[0],
			Accuracy:  similarity,
			Title:     result.Data[0].Data.Source,
			Thumbnail: result.Data[0].Header.Thumbnail,
		},
	)
}

func main() {
	router := gin.Default()

	router.GET("/sauce", getSauce)

	router.Run("localhost:8080")
}
