package configs

import (
	"golang_graphql_postgres/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() *gorm.DB {
	var err error
	dsn := "host=127.0.0.1 user=rey password=rey dbname=golang_graphql port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		log.Fatalf("Could not connect to Postgres: %v", err)
	}

	log.Print("connected to postgres database.")

	err = DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Product{}, &models.Sale{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Print("Tables Created....")
	return DB
}
