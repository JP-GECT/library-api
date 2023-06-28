package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "in search of lost time", Author: "marcel proust", Quantity: 2},
	{ID: "2", Title: "in search of lost time", Author: "marcel proust", Quantity: 5},
	{ID: "3", Title: "in search of lost time", Author: "marcel proust", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")

}
