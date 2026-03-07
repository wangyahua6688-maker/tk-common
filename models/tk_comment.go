package models

import "time"

// WComment 评论表（支持 parent_id 追评）。
type WComment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PostID        uint      `gorm:"index;not null;default:0" json:"post_id"`
	LotteryInfoID uint      `gorm:"index;not null;default:0" json:"lottery_info_id"`
	UserID        uint      `gorm:"index;not null" json:"user_id"`
	ParentID      uint      `gorm:"index;default:0" json:"parent_id"`
	Content       string    `gorm:"size:1000;not null" json:"content"`
	Likes         int64     `gorm:"not null;default:0" json:"likes"`
	Status        int8      `gorm:"not null;default:1" json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (WComment) TableName() string { return "tk_comment" }
