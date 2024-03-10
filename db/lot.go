package db

import (
	"parking/models"
)

func (c *DB) AddLot(lot models.Lot) (*models.Lot, error) {
	var lotDb models.Lot

	qStr := `INSERT INTO lots (name,address,hourly_rate) VALUES ($1,$2,$3) RETURNING *`
	err := c.Pg.QueryRowx(qStr, lot.Name, lot.Address, lot.HourlyRate).StructScan(&lotDb)
	if err != nil {
		return nil, err
	}

	return &lotDb, nil
}

func (c *DB) UpdateLot(id string, lot models.Lot) (*models.Lot, error) {
	var lotDb models.Lot

	qStr := `UPDATE lots SET name = $1, address = $2, hourly_rate = $3 WHERE id = $4 RETURNING *`
	err := c.Pg.QueryRowx(qStr, lot.Name, lot.Address, lot.HourlyRate, id).StructScan(&lotDb)
	if err != nil {
		return nil, err
	}

	return &lotDb, nil
}

func (c *DB) GetLotByID(id string) (*models.Lot, error) {
	var lotDb models.Lot

	qStr := `SELECT * FROM lots WHERE id = $1`
	err := c.Pg.QueryRowx(qStr, id).StructScan(&lotDb)
	if err != nil {
		return nil, err
	}

	return &lotDb, nil
}

func (c *DB) DeleteLotByID(id string) error {

	qStr := `DELETE FROM lots WHERE id = $1`
	_, err := c.Pg.Exec(qStr, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetAllLots() ([]models.Lot, error) {
	allLots := make([]models.Lot, 0)

	qStr := `SELECT * FROM lots`
	res, err := c.Pg.Queryx(qStr)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var lotDb models.Lot
		err := res.StructScan(&lotDb)
		if err != nil {
			return nil, err
		}
		allLots = append(allLots, lotDb)

	}

	return allLots, nil
}

func (c *DB) GetLotBySlotID(slotId string) (*models.Lot, error) {
	var lotDb models.Lot

	qStr := `
        SELECT lots.* 
        FROM lots 
        JOIN slots ON lots.id = slots.lot_id 
        WHERE slots.id = $1
    `
	err := c.Pg.QueryRowx(qStr, slotId).StructScan(&lotDb)
	if err != nil {
		return nil, err
	}

	return &lotDb, nil
}
