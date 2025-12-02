package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

type DefaultModel struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func SetupDatabase() {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DB")

	var err error
	var config gorm.Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbName)

	if os.Getenv("ENABLE_GORM_LOGGER") != "" {
		config = gorm.Config{}
	} else {
		config = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	DB, err = gorm.Open(mysql.Open(dsn), &config)

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}
