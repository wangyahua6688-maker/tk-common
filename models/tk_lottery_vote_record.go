package models

import "time"

// WLotteryVoteRecord 竞猜投票记录表。
// 通过 (lottery_info_id + voter_hash) 唯一约束防止同一设备/指纹重复投票。
type WLotteryVoteRecord struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	LotteryInfoID uint      `gorm:"uniqueIndex:uk_lottery_voter;not null" json:"lottery_info_id"`
	OptionID      uint      `gorm:"not null" json:"option_id"`
	VoterHash     string    `gorm:"size:64;uniqueIndex:uk_lottery_voter;not null" json:"voter_hash"`
	DeviceID      string    `gorm:"size:120" json:"device_id"`
	ClientIP      string    `gorm:"size:64" json:"client_ip"`
	UserAgent     string    `gorm:"size:255" json:"user_agent"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (WLotteryVoteRecord) TableName() string { return "tk_lottery_vote_record" }
