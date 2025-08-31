package dto

import (
	"github.com/eigakan/nats-shared/model"
)

type GetUserRequestDTO struct {
	UserID uint `json:"userId"`
}

type GetUserResponseDTO struct {
	model.User
}
