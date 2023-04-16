package entity

import "time"

type Book struct {
	BookID    int64     `gorm:"primaryKey"`
	BookName  string    `gorm:"not null;type:varchar(100)"`
	Author    string    `gorm:"not null;type:varchar(100)"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
}
