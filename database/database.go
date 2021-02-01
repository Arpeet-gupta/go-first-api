package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// mysql package
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "password"
	hostname = "172.18.56.55:3306"
)

//Db is mysql database connection pool object
// var Db *sql.DB

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

//CreateDataConnectionPool to create new connection pool
func CreateDataConnectionPool(dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)
	db.Close()

	db, err = sql.Open("mysql", dsn(dbName))
	err = db.Ping()
	if err != nil {
		log.Printf("Error %s when creating connections", err)
		return nil, err
	}
	return db, nil
}
