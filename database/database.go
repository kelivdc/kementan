package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func SetupDatabase() {
	godotenv.Load(".env")
	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER1"),
		os.Getenv("DB_PASSWORD1"),
		os.Getenv("DB_HOST1"),
		os.Getenv("DB_NAME1"),
	)

	// dsn2 := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER2"),
	// 	os.Getenv("DB_PASSWORD2"),
	// 	os.Getenv("DB_NAME2"),
	// )

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,        // Don't include params in the SQL log
	// 		Colorful:                  false,       // Disable color
	// 	},
	// )
	db, err := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(2)
	}
	DB = Dbinstance{
		Db: db,
	}
	// db.Use(dbresolver.Register(dbresolver.Config{
	// 	Sources: []gorm.Dialector{mysql.Open(dsn2)},
	// }))
	fmt.Println("== Connected to database ==")
}
