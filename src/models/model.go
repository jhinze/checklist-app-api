package models

import "time"

type Model struct {
	ID        uint       `json:"id",gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty",sql:"index"`
}
