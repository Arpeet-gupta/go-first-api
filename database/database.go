package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

const (
	username = "root"
	password = "password"
	hostname = "172.18.56.55:3306"
)

var (
	//Db contains mysql's database connection pool struct
	Db  *gorm.DB
	err error
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbName)
}

//CreateDataConnectionPool to create new connection pool
func CreateDataConnectionPool(dbName string) error {
	Db, err = gorm.Open(mysql.Open(dsn(dbName)), &gorm.Config{})
	if err != nil {
		log.Printf("Error when creating connection pool: %s", err)
		return err
	}

	Db.Use(prometheus.New(prometheus.Config{
		DBName:          dbName, // use `DBName` as metrics label
		RefreshInterval: 15,     // Refresh metrics interval (default 15 seconds)
		StartServer:     true,   // start http server to expose metrics
		HTTPServerPort:  9090,   // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		}, // user defined metrics
	}))
	return nil
}
