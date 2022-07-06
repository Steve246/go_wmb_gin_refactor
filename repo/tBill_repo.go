package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type BillRepository interface {
	Create(tBil *model.Bill) error
}

type billRepository struct {
	db *gorm.DB
}

func (c *billRepository) Create(tBil *model.Bill) error {
	result := c.db.Create(tBil).Error

	return result
}

func NewBillRepository(db *gorm.DB) BillRepository {
	repo := new(billRepository)
	repo.db = db
	return repo
}
