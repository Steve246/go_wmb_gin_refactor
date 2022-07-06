package model

type Discount struct {
	Id          string    `gorm:"primaryKey" json:"discountId"`
	Description string    `gorm:"size:50" json:"discountDesc"`
	Pct         int       `json:"discountValue"`
	Discount    []*Bill   `gorm:"many2many:customer_discount"`
	BaseModel   BaseModel `gorm:"embedded"`
}

func (d Discount) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_discount"
}
