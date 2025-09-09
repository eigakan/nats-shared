package dto

import (
	"github.com/eigakan/nats-shared/model"
)

type GetUserRequestDTO struct {
	UserID uint64 `json:"userId"`
}

type GetUserResponseDTO struct {
	model.User
}
