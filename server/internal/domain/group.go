package domain

import "time"

type Group struct {
	Id        int       `gorm:"column:id;not null;primaryKey;autoIncrement"`
	OwnerId   string    `gorm:"column:ownerId"`
	Remark    string    `gorm:"column:remark"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	EndedAt   time.Time `gorm:"column:endedAt"`
}
