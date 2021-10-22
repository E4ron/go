package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./html/*.html")
	router.Static("/r/", "./res")
	router.GET("/", handlerIndex)
	router.Run("localhost:8080")
}
func handlerIndex(c *gin.Context) {

	product := []string{
		"Rwgw",
		"egWG",
	}
	c.HTML(200, "index.html", gin.H{
		"Title":   "Hello",
		"Date":    time.Now().String()[0:10],
		"Product": product,
		"A1":      "Главная",
		"A2":      "News",
	})

}
