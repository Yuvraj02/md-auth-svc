package models

import "time"

// User is the site owner / account record.
type User struct {
	ID           string
	Email        string
	Name         string
	DisplayName  string
	Username     string
	Avatar       string
	Bio          string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Analytics is simple denormalized stats for the owner dashboard.
type Analytics struct {
	UserID          string
	TotalViews      int64
	TotalPosts      int
	TotalLikes      int64
	Followers       int
	ViewsThisMonth  int64
	PostsThisMonth  int
	TopPostsJSON    string
}
