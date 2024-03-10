package api

import (
	"fmt"
	"log"
	"parking/db"
	"parking/lib/utils"
	"parking/server/api/validators"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var c *Client

type Client struct {
	Val *validators.FieldValidator
	DB  *db.DB
}

func Init() {
	var err error
	c = &Client{}

	dbUrl := utils.MustOsGetEnv("DB_URL")

	pg := ConnectPg(dbUrl)
	c.DB = db.Init(pg)

	c.Val, err = validators.NewFieldValidator()
	if err != nil {
		log.Panicf("failed to initialize validator: %v", err)
	}
}

func ConnectPg(url string) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to PostgreSQL database!")

	return db
}
