package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sauce struct {
	Database   string  `json:"database"`
	Similarity float32 `json:"similarity"`
	Author     string  `json:"author"`
	Title      string  `json:"title"`
}

var books = []sauce{
	{Database: "Pixiv", Similarity: 94.5, Author: "HEHE", Title: "sample sauce"},
	{Database: "Pixiv", Similarity: 94.5, Author: "HEHE", Title: "sample sauce"},
	{Database: "Pixiv", Similarity: 94.5, Author: "HEHE", Title: "sample sauce"},
	{Database: "Pixiv", Similarity: 94.5, Author: "HEHE", Title: "sample sauce"},
	{Database: "Pixiv", Similarity: 94.5, Author: "HEHE", Title: "sample sauce"},
}

func getSauce(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	router.GET("/sauce", getSauce)

	router.Run("localhost:8080")
}
