package model

type Menu_Price struct {
	Id        string `gorm:"primaryKey" json:"menuPriceId"`
	Menu_Id   Menu
	Price     int       `json:"menuPrice"`
	BaseModel BaseModel `gorm:"embedded"`
}

func (mp Menu_Price) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "m_menu_price"
}
