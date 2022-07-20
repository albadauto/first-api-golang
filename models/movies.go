package models

import (
	"time"

	"gorm.io/gorm"
)

type Movies struct {
	ID        uint           `json:"id" gorm: primaryKey`
	Name      string         `json:"name"`
	Duration  string         `json:"duration"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
