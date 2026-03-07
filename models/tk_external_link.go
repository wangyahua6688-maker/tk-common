package models

import "time"

// WExternalLink 外链配置表。
type WExternalLink struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:80;not null" json:"name"`
	URL       string    `gorm:"size:255;not null" json:"url"`
	Position  string    `gorm:"size:32;index;not null" json:"position"`
	IconURL   string    `gorm:"size:255;not null;default:''" json:"icon_url"`
	GroupKey  string    `gorm:"size:32;index;not null;default:''" json:"group_key"`
	Status    int8      `gorm:"not null;default:1" json:"status"`
	Sort      int       `gorm:"not null;default:0" json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (WExternalLink) TableName() string { return "tk_external_link" }
