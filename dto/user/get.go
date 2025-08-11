package dto

import "time"

type GetUserRequestDTO struct {
	Login  *string `json:"login,omitempty"`
	UserID *uint   `json:"userId,omitempty"`
}

type GetUserResponseDTO struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
