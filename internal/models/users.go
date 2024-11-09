package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null; unique"`
	Email     string    `json:"email" gorm:"not null; unique"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
