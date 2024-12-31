package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var items = []Item{
	{ID: 1, Name: "Item A", Price: 10000, Quantity: 10},
	{ID: 2, Name: "Item B", Price: 20000, Quantity: 5},
}

func main() {
	r := gin.Default()

	r.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)
	})

	r.POST("/checkout", func(c *gin.Context) {
		var cart []Item
		if err := c.ShouldBindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var total float64
		for _, item := range cart {
			total += float64(item.Quantity) * item.Price
		}

		c.JSON(http.StatusOK, gin.H{"total": total})
	})

	r.Run(":8080") // listen on port 8080
}
