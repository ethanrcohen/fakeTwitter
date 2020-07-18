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
		c.Status(500)
		return
	}

	for rows.Next() {
		var p models.Post
		var parentId sql.NullInt32
		err := rows.Scan(&p.Id, &p.Content, &p.UserId, &parentId, &p.Timestamp)
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

type PostForm struct {
	UserId int `form:"user_id"`
	Content string `form:"content"`
	HasParent bool `form:"has_parent"`
	ParentId int  `form:"parent_id"`
}

func MakePost(c *gin.Context) {
	var postForm PostForm
	c.Bind(&postForm)

	var parentIdStr string
	if (postForm.HasParent) {
		parentIdStr = string(postForm.ParentId)
	} else {
		parentIdStr = "NULL"
	}

	queryStr := fmt.Sprintf(
		"INSERT INTO posts (content, user_id, parent_id) VALUES ('%s', %d, %s)",
		postForm.Content,
		postForm.UserId,
		parentIdStr)

	db := db.GetDB()
	db.Begin()
	_, err := db.Exec(queryStr)

	if (err != nil) {
		fmt.Println(err)
		c.Status(500)
	} else {
		c.Status(201)
	}
}