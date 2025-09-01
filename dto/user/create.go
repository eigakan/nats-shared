package dto

type CreateUserRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponseDTO struct {
	Ok bool `json:"ok"`
}
