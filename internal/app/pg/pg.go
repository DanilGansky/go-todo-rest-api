package pg

import (
	"fmt"
	"log"
	"sync"

	"github.com/danikg/go-todo-rest-api/internal/config"
	"github.com/danikg/go-todo-rest-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// GetDB ...
func GetDB(cfg *config.Config) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("host='%s' port=5432 user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)

		var err error
		if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
			log.Fatal("failed to connect db")
		}

		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.TodoList{})
		db.AutoMigrate(&models.TodoItem{})
		db.AutoMigrate(&models.Tag{})
	})
	return db
}
