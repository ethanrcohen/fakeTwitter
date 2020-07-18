package main

import (
	"server/db"
	"server/controllers"
	"time"
 	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const(
	host = "postgres"
	port = 5432
)

// TODO: organize gocode, figure out module structure
func main() {
	// TODO: use retry logic
	time.Sleep(1000 * time.Millisecond)

	db := db.GetDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})

	r.GET("/posts", controllers.GetPosts)
	r.Run()
}
