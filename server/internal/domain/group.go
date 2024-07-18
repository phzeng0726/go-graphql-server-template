package domain

import "time"

type Group struct {
	Id        string     `gorm:"column:groupId"`
	OwnerId   string     `gorm:"column:ownerId"`
	Remark    string     `gorm:"column:remark"`
	CreatedAt time.Time  `gorm:"column:createdAt,serverTimestamp"` // create的時候的timestamp
	EndedAt   time.Time  `gorm:"column:endedAt"`                   // 團購截止日期
	ClosedAt  *time.Time `gorm:"column:closedAt"`                  // 手動關閉的話才填
	Products  []Product  `gorm:"column:products,omitempty"`
}
