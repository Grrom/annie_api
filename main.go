package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	saucenao "github.com/Grrom/annie_api/utils"
)

type sauce struct {
	Link      string  `json:"link"`
	Accuracy  float64 `json:"accuracy"`
	Title     string  `json:"title"`
	Thumbnail string  `json:"thumbnail"`
}

func saucenaoToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error: .Err:%s", err)
	}
	return os.Getenv("SAUCENAO_TOKEN")
}

func getSauce(c *gin.Context) {
	saucenaoClient := saucenao.New(saucenaoToken())
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
		sauce{
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
