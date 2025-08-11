package dto

import "time"

type CheckPasswordRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CheckPasswordResponseDTO struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
