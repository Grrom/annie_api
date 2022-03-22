package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sauce struct {
	Database string  `json:"database"`
	Accuracy float64 `json:"accuracy"`
	Author   string  `json:"author"`
	Title    string  `json:"title"`
}

func getSauce(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sauce{Database: "Pixiv", Accuracy: 94.5, Author: "HEHE", Title: "sample sauce"})
}

func main() {
	router := gin.Default()

	router.GET("/sauce", getSauce)

	router.Run("localhost:8080")
}
