package main

import (
	"errors"
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
	{
		ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Quantity: 10,
	},
	{
		ID: "2", Title: "Golang routers", Author: "Mr. Router", Quantity: 20,
	},
	{
		ID: "3", Title: "Golang concurrency", Author: "Mr. Goroutine", Quantity: 30,
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
func checkoutBook(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Book is out of stock"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}

func createBook(c *gin.Context) {
	var newBook book //var newBook book ifadesi, book türünde yeni bir değişken oluşturur.
	if err := c.BindJSON(&newBook); err != nil {
		return //Eğer JSON verisi geçerli değilse veya bağlama işlemi başarısız olursa, fonksiyon return ifadesi ile sonlanır.
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) //HTTP yanıtı olarak 201 Created durum kodu ve yeni oluşturulan kitabın JSON formatında verisini döner.

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/books/checkout", checkoutBook)
	router.PATCH("/books/return", returnBook)

}
