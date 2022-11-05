package main

import (
	"time"
)

type Blog struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
type Blogs struct {
	AllBlogs []Blog `json:"blogs"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
