package models

import "time"

type AddBookSlotReq struct {
	LotID         string `json:"lot_id" validate:"required,uuid"`
	VehicleNumber string `json:"vehicle_number" validate:"required"`
}

type AddBookSlotResp struct {
	ID            string    `json:"id"`
	SlotID        string    `json:"slot_id"`
	SlotNumber    int       `json:"slot_number"`
	VehicleNumber string    `json:"vehicle_number"`
	CreatedAt     time.Time `json:"created_at"`
	IsParked      bool      `json:"is_parked" db:"is_parked"`
}

type BookSlot struct {
	ID            string     `json:"id" db:"id"`
	SlotID        string     `json:"slot_id" db:"slot_id"`
	VehicleNumber string     `json:"vehicle_number" db:"vehicle_number"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	IsParked      bool       `json:"is_parked" db:"is_parked"`
	EndedAt       *time.Time `json:"ended_at" db:"ended_at"`
	BillAmount    *float64   `json:"bill_amount" db:"bill_amount"`
	SlotNumber    int        `json:"slot_number" db:"slot_number"`
}

type GetBookSlotResp struct {
	ID            string    `json:"id"`
	SlotID        string    `json:"slot_id"`
	VehicleNumber string    `json:"vehicle_number"`
	CreatedAt     time.Time `json:"created_at"`
	IsParked      bool      `json:"is_parked" db:"is_parked"`
	EndedAt       time.Time `json:"ended_at"`
	BillAmount    float64   `json:"bill_amount"`
}

type GetBookStatsResp struct {
	TotalVehicles     int     `json:"total_vehicles"`
	TotalParkingTime  int64   `json:"total_parking_time"`
	TotalFeeCollected float64 `json:"total_fee_collected"`
}

type GetBookStatsDB struct {
	TotalVehicles     int     `json:"total_vehicles" db:"total_vehicles"`
	TotalParkingTime  int64   `json:"total_parking_time" db:"total_parking_time"`
	TotalFeeCollected float64 `json:"total_fee_collected" db:"total_fee_collected"`
}
