package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"reflect"
	"sample-todo/models"
)

func Init() {
	dsn := "host=localhost port=5432 user=potato password=1234 dbname=test sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	if !db.Migrator().HasTable(&models.Task{}) {
		db.AutoMigrate([]models.Task{})
	} else {
		log.Println("table already exists", reflect.TypeOf(models.Task{}).Name())
	}

	DB = db
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
