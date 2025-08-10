package dto

type CheckPasswordRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CheckPasswordResponseDTO struct {
	Valid bool `json:"valid"`
}
