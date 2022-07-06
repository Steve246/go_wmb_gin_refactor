package model

type Bill struct {
	Id string `gorm:"primaryKey" json:"billId"`
	// trans_date    int64       `gorm:"autoCreateTime"`
	Customer_Id []*Discount `gorm:"many2many:customer_discount"` //many to many
	// Table_Id      []Table
	Trans_Type_Id []Trans_Type
	BaseModel     BaseModel `gorm:"embedded"`
}

func (tr Bill) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "t_bill"
}
