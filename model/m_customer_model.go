package model

type Customer struct {
	Id              string    `gorm:"primaryKey" json:"customerId"`
	Customer_Name   string    `json:"customerDesc"`
	Mobile_Phone_No string    `json:"customerPhone"`
	Is_Member       int       `gorm:"default:0" json:"isStatus"`
	BaseModel       BaseModel `gorm:"embedded"`
}

func (c Customer) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_customer"
}
