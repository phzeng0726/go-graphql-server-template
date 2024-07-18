package domain

type Product struct {
	Id    string `gorm:"column:productId"`
	Name  string `gorm:"column:name"`
	Price int    `gorm:"column:price"`
}
