package model

type NatsResponse[T any] struct {
	Status bool   `json:"status"`
	Err    string `json:"err,omitempty"`
	Data   T      `json:"data,omitempty"`
}
