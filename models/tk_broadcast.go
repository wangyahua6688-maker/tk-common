package models

import "time"

// WBroadcast 系统广播表。
type WBroadcast struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:120;not null" json:"title"`
	Content   string    `gorm:"size:500;not null" json:"content"`
	Status    int8      `gorm:"not null;default:1" json:"status"`
	Sort      int       `gorm:"not null;default:0" json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (WBroadcast) TableName() string { return "tk_broadcast" }
