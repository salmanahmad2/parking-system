package db

import (
	"fmt"
	"parking/models"
)

func (c *DB) AddSlot(slot models.Slot) (*models.Slot, error) {
	var slotDb models.Slot

	qStr := `INSERT INTO slots (lot_id,status,number) VALUES ($1,$2,$3) RETURNING *`
	err := c.Pg.QueryRowx(qStr, slot.LotID, slot.Status, slot.Number).StructScan(&slotDb)
	if err != nil {
		return nil, err
	}

	return &slotDb, nil
}

func (c *DB) UpdateSlotStatus(id string, status string) (*models.Slot, error) {
	var slotDb models.Slot

	qStr := `UPDATE slots SET status = $1 WHERE id = $2 RETURNING *`
	err := c.Pg.QueryRowx(qStr, status, id).StructScan(&slotDb)
	if err != nil {
		return nil, err
	}

	return &slotDb, nil
}

func (c *DB) GetSlotByID(id string) (*models.Slot, error) {
	var slotDb models.Slot

	qStr := `SELECT * FROM slots WHERE id = $1`
	err := c.Pg.QueryRowx(qStr, id).StructScan(&slotDb)
	if err != nil {
		return nil, err
	}

	return &slotDb, nil
}

func (c *DB) DeleteSlotByID(id string) error {

	qStr := `DELETE FROM slots WHERE id = $1`
	_, err := c.Pg.Exec(qStr, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetAllSlotsByLotID(lotId, status string) ([]models.Slot, error) {
	allSlots := make([]models.Slot, 0)

	qStr := `SELECT * FROM slots where lot_id =$1`
	if status != "" {
		qStr += fmt.Sprintf(` AND status = '%s'`, status)
	}
	res, err := c.Pg.Queryx(qStr, lotId)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var lotDb models.Slot
		err := res.StructScan(&lotDb)
		if err != nil {
			return nil, err
		}
		allSlots = append(allSlots, lotDb)

	}

	return allSlots, nil
}

func (c *DB) GetLastSlotNumber(lotID string) (int, error) {
	var lastSlotNumber int

	query := `SELECT COALESCE(MAX(number), 0) FROM slots WHERE lot_id = $1`

	err := c.Pg.QueryRow(query, lotID).Scan(&lastSlotNumber)
	if err != nil {
		return 0, err
	}

	return lastSlotNumber, nil
}

func (c *DB) GetNearestSlotByLotID(lotId string) (*models.Slot, error) {
	var slotDb models.Slot

	qStr := `SELECT * FROM slots WHERE lot_id =$1 AND status = 'available' ORDER BY number ASC LIMIT 1`
	err := c.Pg.QueryRowx(qStr, lotId).StructScan(&slotDb)
	if err != nil {
		return nil, err
	}

	return &slotDb, nil
}

func (c *DB) DeleteSlotsByLotID(lotId string) error {

	qStr := `DELETE FROM slots WHERE lot_id = $1`
	_, err := c.Pg.Exec(qStr, lotId)
	if err != nil {
		return err
	}

	return nil
}
