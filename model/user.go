package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Logo      string    `json:"logo,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
