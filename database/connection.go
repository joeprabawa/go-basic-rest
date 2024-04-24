package database

import (
	"fmt"
	"log"
	"os"

	model "github.com/joeprabawa/basic-go-rest/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	fmt.Println(host, user, pass, dbname, port)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, pass, dbname, port)

	dbConect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error connecting to database" + err.Error())
	}

	log.Println("Database connected")
	dbConect.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating database")
	dbConect.AutoMigrate(&model.Destination{})

	Db = dbConect
}
