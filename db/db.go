package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name       string
	Age        int
	Occupation string
}

var DB *gorm.DB

func ConnectDB() {

	db_url := getDBurl()

	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	DB = db
	db.AutoMigrate(&Person{})
	fmt.Println("Connected to DB")

}

func getDBurl() string {
	return viper.Get("DB_URL").(string)
}
