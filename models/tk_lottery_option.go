package models

import "time"

// WLotteryOption 竞猜选项表。
type WLotteryOption struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	LotteryInfoID uint      `gorm:"index;not null" json:"lottery_info_id"`
	OptionName    string    `gorm:"size:32;not null" json:"option_name"`
	Votes         int64     `gorm:"not null;default:0" json:"votes"`
	Sort          int       `gorm:"not null;default:0" json:"sort"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (WLotteryOption) TableName() string { return "tk_lottery_option" }
