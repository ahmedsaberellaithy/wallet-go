package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "wallet-sdk/internal/config"
)

type DB struct {
    *gorm.DB
}

type DBInterface interface {
    Create(value interface{}) *gorm.DB
    First(dest interface{}, conds ...interface{}) *gorm.DB
}

// Global instance
var Instance *DB

func Connect() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        config.AppConfig.DBHost, config.AppConfig.DBUser,
        config.AppConfig.DBPassword, config.AppConfig.DBName, config.AppConfig.DBPort)
    
    gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }
    
    Instance = &DB{DB: gormDB}
}
