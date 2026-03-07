package models

import "time"

// WBanner Banner 配置表：
// - type: ad / official
// - position: home / lottery_detail / post_detail
// 用一张表承载多个展示区与类型配置。
type WBanner struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"size:120;not null" json:"title"`
	ImageURL    string     `gorm:"size:255;not null" json:"image_url"`
	LinkURL     string     `gorm:"size:255" json:"link_url"`
	Type        string     `gorm:"size:32;index;not null" json:"type"`
	Position    string     `gorm:"size:32;index;not null" json:"position"`
	Positions   string     `gorm:"size:255;not null;default:''" json:"positions"`
	JumpType    string     `gorm:"size:20;index;not null;default:'none'" json:"jump_type"`
	JumpPostID  uint       `gorm:"not null;default:0" json:"jump_post_id"`
	JumpURL     string     `gorm:"size:255;not null;default:''" json:"jump_url"`
	ContentHTML string     `gorm:"type:longtext" json:"content_html"`
	Status      int8       `gorm:"not null;default:1" json:"status"`
	Sort        int        `gorm:"not null;default:0" json:"sort"`
	StartAt     *time.Time `json:"start_at"`
	EndAt       *time.Time `json:"end_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (WBanner) TableName() string { return "tk_banner" }
