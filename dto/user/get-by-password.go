package dto

import (
	"github.com/eigakan/nats-shared/model"
)

type GetUserByPasswordRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetUserByPasswordResponseDTO struct {
	model.User
}
