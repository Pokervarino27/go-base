package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	*gorm.DB
}

func NewSqlClient() *PostgresClient {
	DBDriver := os.Getenv("DB_DRIVER")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBSSLMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		DBHost,
		DBPort,
		DBName,
		DBUser,
		DBPassword,
		DBSSLMode,
	)
	db, dbErr := gorm.Open(postgres.New(postgres.Config{
		DriverName: DBDriver,
		DSN:        dsn,
	}))
	if dbErr != nil {
		panic(dbErr)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxIdleTime(30 * time.Second)

	return &PostgresClient{db}
}
