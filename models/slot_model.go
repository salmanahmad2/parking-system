package models

import "time"

type AddSlotReq struct {
	LotID  string `json:"lot_id" validate:"required,uuid"`
	Status string `json:"status" validate:"required,oneof=available unavailable booked"`
}

type SlotResp struct {
	ID        string    `json:"id"`
	LotID     string    `json:"lot_id"`
	Number    int       `json:"number"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Slot struct {
	ID        string    `json:"id" db:"id"`
	LotID     string    `json:"lot_id" db:"lot_id"`
	Number    int       `json:"number" db:"number"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateSlotStatusReq struct {
	Status string `json:"status" validate:"required,oneof=available unavailable booked"`
}
