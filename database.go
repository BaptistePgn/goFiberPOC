package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

type Dbinstance struct {
	Db *gorm.DB
}

var dbInstance Dbinstance

// Connect to the database
func Connect() {
	p := Config("POSTGRES_PORT")
	// convert the port to int
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Error converting port to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Paris", Config("POSTGRES_HOST"), Config("POSTGRES_USER"), Config("POSTGRES_PASSWORD"), Config("POSTGRES_DB"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting to the database")
		// Exit the program
		os.Exit(2)
	}

	fmt.Println("Connected to the database")
	dbInstance = Dbinstance{Db: db}
}
