package dto

import "github.com/eigakan/nats-shared/model"

type CreateUserRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponseDTO struct {
	model.User
}
