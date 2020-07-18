package controllers

import (
	"server/models"
	"server/db"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	db := db.GetDB()
	rows, err := db.Query("SELECT * FROM POSTS;")
	defer rows.Close()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "fail",
		})
	}

	for rows.Next() {
		var p models.Post
		var parentId sql.NullInt32
		err := rows.Scan(&p.Id, &p.Content, &p.UserId, &parentId)
		if err != nil {
			panic(err)
		}

		p.HasParent = parentId.Valid
		if (p.HasParent) {
			p.ParentId = int(parentId.Int32)
		}

		fmt.Printf("%#v", p)
		posts = append(posts, p)
	}

	c.JSON(200, gin.H {
		"posts": posts,
	})
}