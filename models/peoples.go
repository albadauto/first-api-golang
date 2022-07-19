package models

import (
	"time"

	"gorm.io/gorm"
)

type Peoples struct {
	ID        uint           `json:"id" gorm: "primaryKey"`
	Name      string         `json:"name"`
	CPF       string         `json:"cpf"`
	Address   string         `json:"address"`
	Salary    float32        `json:"salary"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
