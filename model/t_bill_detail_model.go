package model

type Bill_Detail struct {
	Id            string `gorm:"primaryKey" json:"billDetailId"`
	Bill_Id       int
	Bill          Bill `gorm:"foreignKey:Bill_Id"`
	Menu_Price_Id int
	Menu_Price    Menu_Price `gorm:"foreignKey:	Menu_Price_Id"`
	Qty           int        `json:"billQty"`
	BaseModel     BaseModel  `gorm:"embedded"`
}

func (tbd Bill_Detail) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "t_bill_detail"
}
