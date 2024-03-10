package models

import "time"

type LotReq struct {
	Name       string  `json:"name" validate:"required"`
	Address    string  `json:"address"`
	HourlyRate float64 `json:"hourly_rate" validate:"min=0"`
}

type LotResp struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	HourlyRate float64   `json:"hourly_rate"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Lot struct {
	ID         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Address    string    `json:"address" db:"address"`
	HourlyRate float64   `json:"hourly_rate" db:"hourly_rate"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
