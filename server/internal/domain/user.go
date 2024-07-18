package domain

type User struct {
	Id   string `gorm:"column:id;not null;primaryKey"`
	Name string `gorm:"column:name"`
}
