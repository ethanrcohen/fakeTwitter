package models

type Post struct {
	Id int
	Content string
	UserId int
	HasParent bool
	ParentId int
}