package models

import "time"

type AddUserReq struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

type UserResp struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        string    `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateUserReq struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone" validate:"required"`
}
