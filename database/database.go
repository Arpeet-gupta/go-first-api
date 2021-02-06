package database

import (
	"fmt"
	"log"

	// mysql package"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	username = "root"
	password = "password"
	hostname = "172.18.56.55:3306"
)

var (
	//Db is mysql database connection pool object
	Db  *gorm.DB
	err error
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, hostname, dbName)
}

//CreateDataConnectionPool to create new connection pool
func CreateDataConnectionPool(dbName string) error {
	// db, err := sql.Open("mysql", dsn(""))
	// if err != nil {
	// 	log.Printf("Error %s when opening DB\n", err)
	// 	return nil, err
	// }

	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelfunc()

	// res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	// if err != nil {
	// 	log.Printf("Error %s when creating DB\n", err)
	// 	return nil, err
	// }

	// no, err := res.RowsAffected()
	// if err != nil {
	// 	log.Printf("Error %s when fetching rows", err)
	// 	return nil, err
	// }
	// log.Printf("rows affected %d\n", no)
	// db.Close()
	//////////////////////////////////////////////////////////////////////////////////////////////////

	Db, err = gorm.Open("mysql", dsn(dbName))
	// err = Db.Ping()
	if err != nil {
		log.Printf("Error %s when creating connections", err)
		return err
	}
	err = Db.DB().Ping()
	if err != nil {
		log.Printf("Error %s when creating connections", err)
		return err
	}
	return nil
}
