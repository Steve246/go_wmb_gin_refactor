package model

type Menu struct {
	Id        string    `gorm:"primaryKey" json:"menuId"`
	Menu_Text string    `gorm:"size:50; not null" json:"menuText"`
	BaseModel BaseModel `gorm:"embedded" `
}

func (m Menu) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_menu"
}
