package dto

type LoginRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}
