package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Sinopse   string         `json:"sinopse"`
	Director  string         `json:"director"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
