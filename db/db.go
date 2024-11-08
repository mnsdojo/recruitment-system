package db

import (
	"log"
	"time"

	"github.com/mnsdojo/recruitment-system/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb(cfg *config.Config) (*gorm.DB, error) {
    // Add retry logic
    var db *gorm.DB
    var err error
    maxRetries := 5

    for i := 0; i < maxRetries; i++ {
        dsn := cfg.GetDBConnectionString()
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("Database connection successfully established.")
            return db, nil
        }

        log.Printf("Failed to connect to database (attempt %d/%d): %v\n", i+1, maxRetries, err)
        time.Sleep(time.Second * 5) // Wait 5 seconds before retrying
    }

    return nil, err
}
