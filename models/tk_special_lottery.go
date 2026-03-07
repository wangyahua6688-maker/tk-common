package models

import "time"

// WSpecialLottery 首页顶部“澳彩/港彩”切换信息表。
type WSpecialLottery struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"size:64;not null" json:"name"`
	Code          string    `gorm:"size:32;uniqueIndex;not null" json:"code"`
	CurrentIssue  string    `gorm:"size:32" json:"current_issue"`
	NextDrawAt    time.Time `json:"next_draw_at"`
	LiveEnabled   int8      `gorm:"not null;default:0" json:"live_enabled"`
	LiveStatus    string    `gorm:"size:16;not null;default:'pending'" json:"live_status"`
	LiveStreamURL string    `gorm:"size:255" json:"live_stream_url"`
	Status        int8      `gorm:"not null;default:1" json:"status"`
	Sort          int       `gorm:"not null;default:0" json:"sort"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (WSpecialLottery) TableName() string { return "tk_special_lottery" }
