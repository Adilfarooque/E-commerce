package db

import (
	"fmt"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
//DB will hold the database connection 
var DB *gorm.DB

func ConnectDatabase(confg config.Config) (*gorm.DB, error) {
	//It constructs the connection string using fmt.Sprintf() with database configuration parameters retrieved from the config.Config struct.
	connectTo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", confg.DBHost, confg.DBUser, confg.DBName, confg.DBPort, confg.DBPassword)
	//Then, it attempts to establish a connection to the PostgreSQL database using gorm.Open(). 
	//It uses the PostgreSQL driver from gorm.io/driver/postgres and passes the connection string along with a default gorm.Config{}.
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	//Handling the error
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database:%w", err)
	}
	DB = db
	return DB, err
}
