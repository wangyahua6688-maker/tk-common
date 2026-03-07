package models

import "time"

// WPostArticle 帖子表（官方发帖+用户发帖）。
type WPostArticle struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	LotteryInfoID uint      `gorm:"index;not null;default:0" json:"lottery_info_id"`
	UserID        uint      `gorm:"index;not null;default:0" json:"user_id"`
	Title         string    `gorm:"size:160;not null" json:"title"`
	CoverImage    string    `gorm:"size:255" json:"cover_image"`
	Content       string    `gorm:"type:text" json:"content"`
	IsOfficial    int8      `gorm:"not null;default:0" json:"is_official"`
	Status        int8      `gorm:"not null;default:1" json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (WPostArticle) TableName() string { return "tk_post_article" }
