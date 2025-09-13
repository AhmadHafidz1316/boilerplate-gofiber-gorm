package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/config"
	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/domain"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	var dial gorm.Dialector
	switch cfg.DBDriver {
	case "postgres":
		dial = postgres.Open(cfg.DBDsn)
	}

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}
	return db, nil
}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Customer{})
}
