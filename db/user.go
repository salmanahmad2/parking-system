package db

import (
	"parking/models"
)

func (c *DB) AddUser(user models.User) (*models.User, error) {
	var userDb models.User

	qStr := `INSERT INTO users (first_name,last_name,phone,email,password,role) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *`
	err := c.Pg.QueryRowx(qStr, user.FirstName, user.LastName, user.Phone, user.Email, user.Password, user.Role).StructScan(&userDb)
	if err != nil {
		return nil, err
	}

	return &userDb, nil
}

func (c *DB) UpdateUser(id string, user models.User) (*models.User, error) {
	var userDb models.User

	qStr := `UPDATE users SET first_name = $1, last_name = $2, phone = $3 WHERE id = $4 RETURNING *`
	err := c.Pg.QueryRowx(qStr, user.FirstName, user.LastName, user.Phone, id).StructScan(&userDb)
	if err != nil {
		return nil, err
	}

	return &userDb, nil
}

func (c *DB) GetUserByID(id string) (*models.User, error) {
	var userDb models.User

	qStr := `SELECT * FROM users WHERE id = $1`
	err := c.Pg.QueryRowx(qStr, id).StructScan(&userDb)
	if err != nil {
		return nil, err
	}

	return &userDb, nil
}

func (c *DB) DeleteUserByID(id string) error {

	qStr := `DELETE FROM users WHERE id = $1`
	_, err := c.Pg.Exec(qStr, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetAllUsers() ([]models.User, error) {
	allUsers := make([]models.User, 0)

	qStr := `SELECT * FROM users`
	res, err := c.Pg.Queryx(qStr)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var userDb models.User
		err := res.StructScan(&userDb)
		if err != nil {
			return nil, err
		}
		allUsers = append(allUsers, userDb)

	}

	return allUsers, nil
}
