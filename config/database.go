package config

import (
	"fmt"
	"go-ecommerce-api/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    user,
    password,
    host,
    port,
    dbname,
)

DB, err = gorm.Open(
	mysql.Open(dsn),
	&gorm.Config{},
)

if err != nil{
	panic(err)
}

err = DB.AutoMigrate(
	&models.Category{},
	&models.Product{},)

if err != nil{
	panic(err)
}

fmt.Println("Connected to MySQL successfully!")

	

}


