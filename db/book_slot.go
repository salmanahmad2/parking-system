package db

import (
	"database/sql"
	"parking/models"
	"time"
)

func (c *DB) AddBookSlot(book models.BookSlot) (*models.BookSlot, error) {
	var bookDb models.BookSlot

	qStr := `INSERT INTO book_slot (slot_id,vehicle_number) VALUES ($1,$2) RETURNING *`
	err := c.Pg.QueryRowx(qStr, book.SlotID, book.VehicleNumber).StructScan(&bookDb)
	if err != nil {
		return nil, err
	}

	return &bookDb, nil
}

func (c *DB) UpdateBookSlot(id string, book models.BookSlot) (*models.BookSlot, error) {
	var bookDb models.BookSlot

	qStr := `UPDATE book_slot SET ended_at = $1, bill_amount = $2, is_parked = $3 WHERE id = $4 RETURNING *`
	err := c.Pg.QueryRowx(qStr, book.EndedAt, book.BillAmount, book.IsParked, id).StructScan(&bookDb)
	if err != nil {
		return nil, err
	}

	return &bookDb, nil
}

func (c *DB) GetBookSlotByID(id string) (*models.BookSlot, error) {
	var bookDb models.BookSlot

	qStr := `SELECT * FROM book_slot WHERE id = $1`
	err := c.Pg.QueryRowx(qStr, id).StructScan(&bookDb)
	if err != nil {
		return nil, err
	}

	return &bookDb, nil
}

func (c *DB) DeleteBookSlotByID(id string) error {

	qStr := `DELETE FROM book_slot WHERE id = $1`
	_, err := c.Pg.Exec(qStr, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetBookedSlotsStats(startTime, endTime time.Time) (*models.GetBookStatsDB, error) {
	var stats models.GetBookStatsDB
	qStr := `
		SELECT COUNT(id) AS total_vehicles,
		       COALESCE(SUM(EXTRACT(EPOCH FROM ended_at - created_at)::integer), 0) AS total_parking_time,
		       COALESCE(SUM(bill_amount), 0) AS total_fee_collected
		FROM book_slot
		WHERE created_at >= $1 AND created_at < $2 AND is_parked=false
	`
	err := c.Pg.QueryRowx(qStr, startTime, endTime).StructScan(&stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (c *DB) DeleteBookedSlotsByLotID(lotId string) error {

	qStr := `DELETE FROM book_slot WHERE slot_id IN (SELECT id FROM slots WHERE lot_id = $1)`
	_, err := c.Pg.Exec(qStr, lotId)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) DeleteBookedSlotsBySlotID(slotId string) error {

	qStr := `DELETE FROM book_slot WHERE slot_id = $1`
	_, err := c.Pg.Exec(qStr, slotId)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) IsSlotBooked(slotId string) (bool, error) {
	var id *string
	qStr := `SELECT id FROM book_slot WHERE slot_id = $1 AND is_parked=true`
	err := c.Pg.QueryRowx(qStr, slotId).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if id != nil {
		return true, nil
	}

	return false, nil
}

func (c *DB) GetAllBookSlotsByLotID(lotId string) ([]models.BookSlot, error) {
	allBookSlots := make([]models.BookSlot, 0)

	qStr := `SELECT bs.*,s.number AS slot_number FROM book_slot bs 
	INNER JOIN slots s ON bs.slot_id = s.id
	where s.lot_id =$1 AND bs.is_parked=true`

	res, err := c.Pg.Queryx(qStr, lotId)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var bookSlotDb models.BookSlot
		err := res.StructScan(&bookSlotDb)
		if err != nil {
			return nil, err
		}
		allBookSlots = append(allBookSlots, bookSlotDb)

	}

	return allBookSlots, nil
}
