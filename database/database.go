package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func SetupDatabase() {
	godotenv.Load(".env")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME_1"),
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:               newLogger,
		DisableAutomaticPing: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(2)
	}
	DB = Dbinstance{
		Db: db,
	}
	fmt.Println("== Connected to database ==")
}
