package database

import (
	"context"
	"database/sql"
	"log"
	"time"
)

//Createtable will check if table is exist or not, if not it will table.
func Createtable(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	query := "CREATE TABLE IF NOT EXISTS posts(title text, body text, author varchar(200))"

	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating posts table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating posts table: %d", rows)
	return nil
}
