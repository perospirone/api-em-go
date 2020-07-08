package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Conta struct {
	ID uint `gorm:"primary_key"`
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&Conta{})
}
