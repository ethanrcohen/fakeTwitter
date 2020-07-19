package models

import "time"

type Post struct {
	Id        int
	Content   string
	UserId    int
	HasParent bool
	ParentId  int
	Timestamp time.Time
}
