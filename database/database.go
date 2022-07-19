package database

import (
	"log"

	"github.com/albadauto/projeto/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=5432 user=postgres dbname=lafontana sslmode=disable password=root"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil{
		log.Fatal("error:", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	migrations.RunMigrations(db)
}


func CloseDB() error{
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()

	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB{
	return db
}
