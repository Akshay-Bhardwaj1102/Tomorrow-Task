package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task.com/model"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=postgres password=manager dbname=Leave port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.EmpLeave{}, &model.File{})

	DB = db
}
