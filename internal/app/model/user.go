package model

import (
	"time"
)

type User struct {
	Id             int       `json:"id"`
	Email          string    `json:"email"`
	RoleName       string    `json:"roleName"`
	HashedPassword string    `json:"-"`
	AccountNumber  int64     `json:"accountNumber"`
	Balance        int64     `json:"balance"`
	CreatedAt      time.Time `json:"createdAt"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
