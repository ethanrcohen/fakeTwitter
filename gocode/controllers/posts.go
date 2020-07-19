package controllers

import (
	"database/sql"
	"fmt"

	"github.com/ethanrcohen/fakeTwitter/db"
	"github.com/ethanrcohen/fakeTwitter/models"
	"github.com/gin-gonic/gin"
)

// GetPosts fetches all posts.
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
		var parentID sql.NullInt32
		err := rows.Scan(&p.ID, &p.Content, &p.UserID, &parentID, &p.Timestamp)
		if err != nil {
			panic(err)
		}

		p.HasParent = parentID.Valid
		if p.HasParent {
			p.ParentID = int(parentID.Int32)
		}

		fmt.Printf("%#v", p)
		posts = append(posts, p)
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

type postForm struct {
	UserID    int    `form:"user_id"`
	Content   string `form:"content"`
	HasParent bool   `form:"has_parent"`
	ParentID  int    `form:"parent_id"`
}

// MakePost creates a post from the provided form data.
func MakePost(c *gin.Context) {
	var postForm postForm
	c.Bind(&postForm)

	var parentIDStr string
	if postForm.HasParent {
		parentIDStr = string(postForm.ParentID)
	} else {
		parentIDStr = "NULL"
	}

	queryStr := fmt.Sprintf(
		"INSERT INTO posts (content, user_id, parent_id) VALUES ('%s', %d, %s)",
		postForm.Content,
		postForm.UserID,
		parentIDStr)

	db := db.GetDB()
	db.Begin()
	_, err := db.Exec(queryStr)

	if err != nil {
		fmt.Println(err)
		c.Status(500)
	} else {
		c.Status(201)
	}
}
