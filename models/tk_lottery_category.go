package models

import "time"

// WLotteryCategory 图库分类配置表。
// 用途：
// 1. 首页分类条按排序展示；
// 2. “更多分类”页可搜索；
// 3. 支持后台启停与是否首页展示控制。
type WLotteryCategory struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CategoryKey    string    `gorm:"column:category_key;size:32;uniqueIndex;not null" json:"category_key"`
	Name           string    `gorm:"size:32;not null" json:"name"`
	SearchKeywords string    `gorm:"size:255;not null;default:''" json:"search_keywords"`
	ShowOnHome     int8      `gorm:"not null;default:1" json:"show_on_home"`
	Status         int8      `gorm:"not null;default:1" json:"status"`
	Sort           int       `gorm:"not null;default:0" json:"sort"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (WLotteryCategory) TableName() string { return "tk_lottery_category" }
