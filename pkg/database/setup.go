package database

import (
	"fmt"
	"log"

	"github.com/irhendra09/chat-app/app/models"
	"github.com/irhendra09/chat-app/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgres() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.GetEnv("POSTGRES_HOST", "localhost"),
		env.GetEnv("POSTGRES_USER", "appuser"),
		env.GetEnv("POSTGRES_PASSWORD", "secretpassword"),
		env.GetEnv("POSTGRES_DB", "messengerdb"),
		env.GetEnv("POSTGRES_PORT", "5432"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&models.User{}, &models.UserSession{})
	if err != nil {
		log.Fatal("Database migration failed: ", err)
	}
	log.Println("Database migration succeeded")
}
