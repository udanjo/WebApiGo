package database

import (
	"log"
	"time"

	"github.com/ueverson/WebApiGo/database/migrations"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	//connString := "host=localhost port=5434 user=sa dbname=movieDB sslmode=disable password=udanjo2011@"
	connString := "sqlserver://sa:udanjo2011@@localhost:5434?database=movieDB"
	//sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm

	database, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("error database: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
