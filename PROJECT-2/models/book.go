package models

import "time"

type Book struct {
	Id         uint   `gorm:"primaryKey"`
	Name_book  string `gorm:"not null;type:varchar(100)"`
	Author     string `gorm:"not null;type:varchar(100)"`
	Created_at time.Time
	Updated_at time.Time
}
