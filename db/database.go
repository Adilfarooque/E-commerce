package db

import (
	"fmt"

	"github.com/Adilfarooque/Footgo/config"
	"github.com/Adilfarooque/Footgo/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config config.Config) (*gorm.DB, error) {
    connectTo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
    db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    DB = db
    db.AutoMigrate(&domain.Admin{})
    db.AutoMigrate(&domain.User{})
    return DB, nil
}
