package main

import (
	"time"

	"github.com/ethanrcohen/fakeTwitter/controllers"
	"github.com/ethanrcohen/fakeTwitter/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host = "postgres"
	port = 5432
)

// TODO: organize gocode, figure out module structure
func main() {
	// TODO: use retry logic
	time.Sleep(1000 * time.Millisecond)

	db := db.GetDB()
	defer db.Close()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/post", controllers.MakePost)
	router.GET("/posts", controllers.GetPosts)
	router.Run(":8080")
}
