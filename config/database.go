package config

import (
	"fmt";
	"log";
	"os";
	"gorm.io/gorm";
	"gorm.io/driver/postgres";
)

var DB *gorm.DB;

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"),
)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB connection failed!");
	}

	DB = db;

	log.Println("Database connected!");
}