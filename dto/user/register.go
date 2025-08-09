package dto

type RegisterRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponseDTO struct {
	Ok bool `json:"ok"`
}
