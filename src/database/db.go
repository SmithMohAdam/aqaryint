package database

import (
	"aqaryint/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	dsn := "host=localhost user=postgres password=smith dbname=postgres port=5432 sslmode=disable "
	//dsn := fmt.Sprint(dbHost, username, password, dbName, dbPort, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Student{}) //Department
	//db.AutoMigrate(&models.Department{})
	DB = db
	return nil
}
