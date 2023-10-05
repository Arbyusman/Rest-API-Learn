package model

import "time"

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type HttpResponse struct {
	Data       interface{} `json:"data"`
	MetaData   MetaData    `json:"metadata"`
	Pagination *Pagination `json:"pagination"`
}

type MetaData struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type AuthResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID        string    `json:"id" `
	Name      string    `json:"name" `
	Email     string    `json:"email" `
	Balance   float64   `json:"balance" `
	Pin       string    `json:"pin" `
	Phone     string    `json:"phone"`
	Image     string    `json:"image"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
