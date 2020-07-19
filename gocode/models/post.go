package models

import "time"

// Post represents a single user content upload, or post, on fakeTwitter.
type Post struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	HasParent bool      `json:"has_parent"`
	ParentID  int       `json:"parent_id"`
	Timestamp time.Time `json:"timestamp"`
}
